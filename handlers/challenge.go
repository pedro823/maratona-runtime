package handlers

import (
	"log"
	"net/http"

	"github.com/go-pg/pg/v9"
	"github.com/pedro823/maratona-runtime/handlers/responses"
	"github.com/pedro823/maratona-runtime/model"
	"github.com/pedro823/maratona-runtime/util"
)

func GetAllChallenges(req *http.Request, res *util.JSONRenderer, logger *log.Logger, db *pg.DB) {
	var challenges []model.Challenge
	err := db.Model(&challenges).Select()
	if err != nil {
		panic(err)
	}
	res.JSON(200, responses.AllChallengesResponse{Challenges: challenges})
}

func UploadChallenge(req *http.Request, res http.ResponseWriter, logger *log.Logger, db *pg.DB) {

}
