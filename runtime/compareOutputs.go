package runtime

import (
	"fmt"
	"github.com/pedro823/maratona-runtime/model"
)

func compareOutputs(expected, actual []byte) model.ChallengeResult {
	lastNewLinePosition := 0
	lineNo := 0
	for i := range smallestBetween(expected, actual) {
		a, b := expected[i], actual[i]
		if a != b {
			nextNewLineExpected := findNextNewLine(expected, i)
			nextNewLineActual := findNextNewLine(actual, i)
			return wrongAnswer(lineNo, expected[lastNewLinePosition:nextNewLineExpected], actual[lastNewLinePosition:nextNewLineActual])
		}
		if a == byte('\n') {
			lastNewLinePosition = i
			lineNo++
		}
	}

	return model.ChallengeResult{
		Status: model.Success,
		Reason: "Correct Answer",
	}
}

func wrongAnswer(line int, expected []byte, actual []byte) model.ChallengeResult {
	reason := fmt.Sprintf("Wrong answer on line %d: expected %s, got %s", line, string(expected), string(actual))
	return model.ChallengeResult{Status: model.WrongAnswer, Reason: reason}
}

func findNextNewLine(v []byte, from int) int {
	for i, letter := range v[from:] {
		if letter == byte('\n') {
			return from + i
		}
	}
	return -1
}

func smallestBetween(a, b []byte) []byte {
	if len(a) < len(b) {
		return a
	}
	return b
}
