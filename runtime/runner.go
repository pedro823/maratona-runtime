package runtime

import (
	"context"
	"fmt"
	"github.com/pedro823/maratona-runtime/model"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

func CompileAndRun(program []byte, challenge model.Challenge, compiler AvailableCompiler, resultOut chan<- model.ChallengeResult) {
	executable, tempFiles, err := compiler.Compile(program)
	//noinspection GoNilness
	defer tempFiles.Clean()
	if err != nil {
		result := compilerError(err)
		resultOut <- result
		return
	}

	input, expectedOutput := challenge.Input.RawData, challenge.Output.RawData

	inputFile, err := createInputFile(input)
	if err != nil {
		result := compilerError(err)
		resultOut <- result
		return
	}
	defer os.Remove(inputFile.Name())

	actualOutput := make(chan []byte)
	errorOutput := make(chan error)

	ctx, cancelFunc := timerContext(challenge)
	defer cancelFunc()
	go execute(ctx, executable, inputFile, actualOutput, errorOutput)

	select {
	case <-ctx.Done():
		result := timeLimitExceeded(challenge.Timeout)
		resultOut <- result
		return
	case err = <-errorOutput:
		result := model.ChallengeResult{Status: model.RuntimeError, Reason: err.Error()}
		resultOut <- result
		return
	case out := <-actualOutput:
		result := compareOutputs(expectedOutput, out)
		resultOut <- result
		return
	}
}

func timeLimitExceeded(timeout time.Duration) model.ChallengeResult {
	reason := fmt.Sprintf("Exceeded time limit of %v seconds", timeout.Seconds())
	return model.ChallengeResult{Status: model.TimeLimitExceeded, Reason: reason}
}

func execute(ctx context.Context, executable []string, inputFile *os.File, output chan<- []byte, errorOutput chan<- error) {
	cmd := exec.CommandContext(ctx, executable[0], executable[1:]...)
	cmd.Stdin = inputFile
	programOutput, err := cmd.Output()
	if err != nil {
		errorOutput <- err
		return
	}
	output <- programOutput
}

func createInputFile(input []byte) (*os.File, error) {
	file, err := ioutil.TempFile("", "input")
	if err != nil {
		return file, err
	}
	_, err = file.Write(input)
	return file, err
}

func compilerError(err error) model.ChallengeResult {
	return model.ChallengeResult{Status: model.CompilerError, Reason: err.Error()}
}

func timerContext(challenge model.Challenge) (context.Context, context.CancelFunc) {
	return context.WithDeadline(context.Background(), time.Now().Add(challenge.Timeout))
}
