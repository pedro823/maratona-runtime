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
	timeLimitExceededDir = "testData/timeLimitExceeded"
	wrongAnswerDir = "testData/wrongAnswer"
)

func timeTrack(t *testing.T, start time.Time, name string) {
	elapsed := time.Since(start)
	t.Logf("%s: %s", name, elapsed)
}

func TestProgramOutput(t *testing.T) {
	t.Parallel()

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

	defer timeTrack(t, time.Now(), "TestProgramOutput")
	go CompileAndRun(challengeAttempt, challenge, &Python3{}, resultChan)

	result := <-resultChan
	if result.Status != model.Success {
		t.Fatalf("Expected result to be success, got result %v", result)
	}
}

func TestRuntimeSuccess(t *testing.T) {
	t.Parallel()

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

	defer timeTrack(t, time.Now(), "TestRuntimeSuccess")
	go CompileAndRun(challengeAttempt, challenge, &CPlusPlus11{}, resultChan)

	result := <-resultChan
	if result.Status != model.Success {
		t.Fatalf("Expected result to be success, got result %v", result)
	}
}

func TestTimeLimitExceeded(t *testing.T) {
	t.Parallel()

	challengeAttempt := readFromFile(t, timeLimitExceededDir, "attempt.go")
	challenge := model.Challenge{
		Timeout:     1 * time.Second,
		Input:       &model.ChallengeInput{},
		Output:      &model.ChallengeOutput{},
	}
	resultChan := make(chan model.ChallengeResult)

	defer timeTrack(t, time.Now(), "TestTimeLimitExceeded")
	go CompileAndRun(challengeAttempt, challenge, &Go{}, resultChan)

	result := <-resultChan
	if result.Status != model.TimeLimitExceeded {
		t.Fatalf("Expected result to be a Time Limit Exceeded, got %v", result)
	}
}

func TestWrongAnswer(t *testing.T) {
	t.Parallel()

	challengeAttempt := readFromFile(t, wrongAnswerDir, "attempt.py")
	input := readFromFile(t, wrongAnswerDir, "input.txt")
	output := readFromFile(t, wrongAnswerDir, "output.txt")

	challenge := model.Challenge{
		Timeout: 2 * time.Second,
		Input: &model.ChallengeInput{RawData:input},
		Output: &model.ChallengeOutput{RawData:output},
	}

	resultChan := make(chan model.ChallengeResult)

	defer timeTrack(t, time.Now(), "TestWrongAnswer")
	go CompileAndRun(challengeAttempt, challenge, &Python3{}, resultChan)

	result := <-resultChan
	if result.Status != model.WrongAnswer {
		t.Fatalf("Expected result to be a Wrong Answer, got %v", result)
	}
	t.Logf("In Wrong answer: %s", result.Reason)
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


