package runtime

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

type AvailableCompiler interface {
	Compile(program []byte) ([]string, TempFiles, error)
}

type TempFiles []string

func (t TempFiles) Clean() {
	for _, file := range t {
		_ = os.Remove(file)
	}
}

type Python3 struct{}

func (c *Python3) Compile(program []byte) ([]string, TempFiles, error) {
	file, err := ioutil.TempFile("", "attempt*.py")
	if err != nil {
		return nil, nil, err
	}

	tempFiles := TempFiles{file.Name()}

	_, err = file.Write(program)
	if err != nil {
		return nil, tempFiles, err
	}

	return []string{"python3", file.Name()}, tempFiles, nil
}

type CPlusPlus11 struct{}

func (c *CPlusPlus11) Compile(program []byte) ([]string, TempFiles, error) {
	file, err := ioutil.TempFile("", "attempt*.cc")
	if err != nil {
		return nil, nil, err
	}

	tempFiles := make(TempFiles, 0, 2)
	tempFiles = append(tempFiles, file.Name())

	_, err = file.Write(program)
	if err != nil {
		return nil, tempFiles, err
	}

	executableName := fmt.Sprintf("%s.exe", file.Name())
	cmd := exec.Command("g++", "-std=c++11", file.Name(), "-o", executableName)
	output, err := cmd.Output()
	if err != nil {
		return nil, nil, errors.New(string(output))
	}

	tempFiles = append(tempFiles, executableName)

	return []string{executableName}, tempFiles, nil
}

type Java8 struct{}

//func (c *Java8) Compile(programFile *os.File) ([]string, error) {
//
//}
