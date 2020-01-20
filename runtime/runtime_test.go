package runtime

import (
	"github.com/pedro823/maratona-runtime/model"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

const (
	runtimeSuccessDir = "testData/runtimeSuccess"
	programOutputDir = "testData/programOutput"
)

func TestProgramOutput(t *testing.T) {
	outputBytes := readFromFile(t, programOutputDir, "output.txt")
	challengeAttempt := readFromFile(t, programOutputDir, "attempt.py")

	input := model.ChallengeInput{}
	output := model.ChallengeOutput{RawData: outputBytes}
	challenge := model.Challenge{
		Timeout: 2 * time.Second,
		Input: &input,
		Output: &output,
	}

	resultChan := make(chan model.ChallengeResult)
	go CompileAndRun(challengeAttempt, challenge, &Python3{}, resultChan)

	result := <-resultChan
	if result.Status != model.Success {
		t.Fatalf("Expected result to be success, got result %v", result)
	}
}

func TestRuntimeSuccess(t *testing.T) {
	inputBytes := readFromFile(t, runtimeSuccessDir, "input.txt")
	outputBytes := readFromFile(t, runtimeSuccessDir, "output.txt")
	challengeAttempt := readFromFile(t, runtimeSuccessDir, "attempt.cc")

	input := model.ChallengeInput{RawData: inputBytes}
	output := model.ChallengeOutput{RawData:outputBytes}
	challenge := model.Challenge{
		Timeout: 2 * time.Second,
		Input: &input,
		Output: &output,
	}

	resultChan := make(chan model.ChallengeResult)

	go CompileAndRun(challengeAttempt, challenge, &CPlusPlus11{}, resultChan)

	result := <-resultChan
	if result.Status != model.Success {
		t.Fatalf("Expected result to be success, got result %v", result)
	}
}

func TestTimeLimitExceeded(t *testing.T) {

}

func readFromFile(t *testing.T, directory, file string) []byte {
	fp, err := os.Open(directory + "/" + file)
	if err != nil {
		t.Fatal(err)
	}

	content, err := ioutil.ReadAll(fp)
	if err != nil {
		t.Fatal(err)
	}

	return content
}


