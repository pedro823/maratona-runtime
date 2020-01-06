package handlers

import (
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

func UploadChallenge(req *http.Request, res http.ResponseWriter, logger *util.TimeLogger, db *pg.DB) {

}
