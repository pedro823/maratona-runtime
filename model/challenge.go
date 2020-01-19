package model

import "time"

type Challenge struct {
	ID          int64            `json:"id",pg:"pk"`
	Title       string           `json:"title,notnull"`
	Description string           `json:"description,notnull"`
	Timeout     time.Duration    `json:"-"`
	Input       *ChallengeInput  `json:"-"`
	Output      *ChallengeOutput `json:"-"`
}

type ChallengeInput struct {
	ID      int64 `pg:"pk"`
	RawData []byte
}

type ChallengeOutput struct {
	ID      int64 `pg:"pk"`
	RawData []byte
}
