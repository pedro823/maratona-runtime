package handlers

import (
	"github.com/go-pg/pg"
	"github.com/pedro823/maratona-runtime/errors"
	"github.com/pedro823/maratona-runtime/model"
	"github.com/pedro823/maratona-runtime/runtime"
	"github.com/pedro823/maratona-runtime/util"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	challengeParam       = "challenge"
	errorBackoffInterval = 5 * time.Minute
	backOffMaxAttempts   = 5
)

func ExtractAndExecute(req *http.Request, params map[string]string, res *util.JSONRenderer, logger *util.TimeLogger, db *pg.DB, context util.ContextMap) {
	// Extract
	challengeId, ok := params[challengeParam]
	if !ok {
		errors.NewHTTPError(400, "No challenge was set in request").WriteJSON(res)
		return
	}

	id, err := strconv.Atoi(challengeId)
	if err != nil {
		errors.NewHTTPError(400, "Expected challenge to be an integer").WriteJSON(res)
		return
	}

	user := context[util.UserContextKey].(*model.User)
	compiler := context[util.CompilerContextKey].(runtime.AvailableCompiler)

	challenge := model.Challenge{ID: int64(id)}
	err = db.Model(&challenge).Select()
	if err != nil {
		errors.NewHTTPError(500, "Internal error: "+err.Error()).WriteJSON(res)
		return
	}

	if challenge.Input == nil || challenge.Output == nil {
		// not found or not set
		errors.NewHTTPError(404, "Challenge not found").WriteJSON(res)
		return
	}

	program, err := ioutil.ReadAll(req.Body)
	if err != nil {
		errors.NewHTTPError(400, "error reading request body").WriteJSON(res)
		return
	}

	attempt, err := createChallengeAttempt(db, challenge, user)
	if err != nil {
		logger.TimePrintf("[ERROR]: Database returned error %v", err)
		errors.NewHTTPError(500, "Unexpected error. Please contact an admin")
		return
	}

	// Run
	go watchRuntime(program, compiler, attempt, db, logger)

	res.JSON(201, map[string]string{"attemptId": attempt.Hash})
}

func createChallengeAttempt(db *pg.DB, challenge model.Challenge, user *model.User) (model.ChallengeAttempt, error) {
	attempt := model.ChallengeAttempt{
		Hash:      util.CreateHash(),
		User:      user,
		Challenge: &challenge,
		Result: &model.ChallengeResult{
			Status: model.InProgress,
		},
	}
	err := db.Insert(&attempt)
	return attempt, err
}

func watchRuntime(program []byte, compiler runtime.AvailableCompiler, attempt model.ChallengeAttempt, db *pg.DB, logger *util.TimeLogger) {
	resultChan := make(chan model.ChallengeResult)
	go runtime.CompileAndRun(program, *attempt.Challenge, compiler, resultChan)
	logger.TimePrintf("Started attempt of challenge %s for user %s", attempt.Challenge.Title, attempt.User.CTFID)

	result := <-resultChan
	logger.TimePrintf("Result of attempt [challenge=%s user=%s]: %s", attempt.Challenge.Title, attempt.User.CTFID, result.Reason)

	attempt.Result = &result
	for i := 0; i < 5; i++ {
		err := db.Insert(&attempt)
		if err != nil {
			logger.TimePrintf("[ERROR]: Could not insert result %s of attempt %v into database. Retry %d of %d", result.Reason, attempt, i, backOffMaxAttempts)
		}
		time.Sleep(errorBackoffInterval)
	}
}
