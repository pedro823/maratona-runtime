package responses

import (
	"github.com/pedro823/maratona-runtime/model"
)

type AllChallengesResponse struct {
	Challenges []model.Challenge `json:"challenges"`
}
