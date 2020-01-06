package requests

type CreateChallengeRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
