package model

var allModels = []interface{}{
	(*Challenge)(nil),
	(*ChallengeInput)(nil),
	(*ChallengeOutput)(nil),
	(*User)(nil),
	(*ChallengeAttempt)(nil),
	(*ChallengeResult)(nil),
}

func GetAll() []interface{} {
	return allModels
}
