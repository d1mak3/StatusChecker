package utils

import (
	"io"
	"strings"
)

func GetReadStreamLine(reader io.Reader) func() (string, error) {
	lineBuffer := ""

	return func() (string, error) {
		buffer := make([]byte, 16)

		indexOfNewLine := strings.Index(lineBuffer, "\n")
		for indexOfNewLine == -1 {
			_, err := reader.Read(buffer)
			if err != nil {
				return lineBuffer, err
			}

			readString := string(buffer)
			lineBuffer += readString
			indexOfNewLine = strings.Index(readString, "\n")
		}

		if indexOfNewLine == -1 {
			return lineBuffer, nil
		}

		result := lineBuffer[:indexOfNewLine]
		lineBuffer = lineBuffer[indexOfNewLine+1:] // remove the \n
		return result, nil
	}
}
