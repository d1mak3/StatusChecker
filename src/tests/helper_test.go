package testing

import (
	"io"
	"os"
	"testing"
	"utils"
)

func Test_ReadStreamByLines_ShouldBeOk(t *testing.T) {
	file, err := os.Open(testFile)
	if err != nil {
		t.Error(err.Error())
	}

	readLine := utils.GetReadStreamLine(file)
	lines := make([]string, 0)
	for err == nil {
		line := ""
		line, err = readLine()
		if line != "" {
			lines = append(lines, line)
		}
	}

	if err != io.EOF {
		t.Error(err.Error())
	}

	if len(lines) < linesInTestFile {
		t.Errorf("expected %d lines but found %d", linesInTestFile, len(lines))
	}
}

func Test_ReadEmptyStreamByLines_ShouldBeOk(t *testing.T) {
	file, err := os.Open(emptyFile)
	if err != nil {
		t.Error(err.Error())
	}

	readLine := utils.GetReadStreamLine(file)
	lines := make([]string, 0)
	for err == nil {
		line := ""
		line, err = readLine()
		if line != "" {
			lines = append(lines, line)
		}
	}

	if err != io.EOF {
		t.Error(err.Error())
	}

	if len(lines) > 0 {
		t.Errorf("expected %d lines but found %d", 0, len(lines))
	}
}
