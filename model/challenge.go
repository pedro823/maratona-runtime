package model

type Challenge struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Input       *ChallengeInput
	Output      *ChallengeOutput
}

type ChallengeInput struct {
	ID      int64
	RawData string
}

type ChallengeOutput struct {
	ID      int64
	RawData string
}
