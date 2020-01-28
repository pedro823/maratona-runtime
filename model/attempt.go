package model

type ResultStatus int

const (
	InProgress ResultStatus = iota
	Success
	WrongAnswer
	TimeLimitExceeded
	MemoryLimitExceeded
	CompilerError
	RuntimeError
)

type ChallengeAttempt struct {
	Hash      string `pg:"pk"`
	User      *User
	Challenge *Challenge
	Result    *ChallengeResult
}

type ChallengeResult struct {
	Status ResultStatus
	Reason string
}
