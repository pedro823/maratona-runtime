package model

type ResultStatus int

const (
	InProgress ResultStatus = iota
	Success
	WrongAnswer
	TimeLimitExceeded
	MemoryLimitExceeded
	CompilerError
)

type ChallengeAttempt struct {
	ID        int64
	User      *User
	Challenge *Challenge
	Result    *ChallengeResult
}

type ChallengeResult struct {
	Status ResultStatus
	Reason string
}
