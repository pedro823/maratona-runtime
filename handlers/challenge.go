package handlers

import (
	"encoding/json"
	"github.com/pedro823/maratona-runtime/errors"
	"github.com/pedro823/maratona-runtime/handlers/requests"
	"io/ioutil"
	"net/http"

	"github.com/go-pg/pg/v9"
	"github.com/pedro823/maratona-runtime/handlers/responses"
	"github.com/pedro823/maratona-runtime/model"
	"github.com/pedro823/maratona-runtime/util"
)

func GetAllChallenges(req *http.Request, res *util.JSONRenderer, logger *util.TimeLogger, db *pg.DB) {
	defer logger.TimePrintf("Challenge database was accessed with admin privileges")

	var challenges []model.Challenge
	err := db.Model(&challenges).Select()
	if err != nil {
		panic(err)
	}

	res.JSON(200, responses.AllChallengesResponse{Challenges: challenges})
}

func UploadChallenge(req *http.Request, res *util.JSONRenderer, logger *util.TimeLogger, db *pg.DB) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		defer logger.TimePrintf("Error reading request body: %v", err)
		errors.NewHTTPError(http.StatusBadRequest, "Could not read request body").WriteJSON(res)
		return
	}
	var formattedRequest requests.CreateChallengeRequest

	err = json.Unmarshal(body, &formattedRequest)
	if err != nil {
		defer logger.TimePrintf("Malformed request to UploadChallenge: %v", err)
		errors.NewHTTPError(http.StatusBadRequest, err.Error()).WriteJSON(res)
		return
	}

	challenge := &model.Challenge{Title: formattedRequest.Title, Description: formattedRequest.Description}
	err = db.Insert(challenge)
	if err != nil {
		defer logger.TimePrintf("Could not insert new challenge into database. Service unavailable? Error: %v", err)
		errors.NewHTTPError(http.StatusInternalServerError, "Could not write challenge to database").WriteJSON(res)
		return
	}

	res.JSON(201, nil)
}
