package utils

type StreamReadingError struct {
	message string
}

func GetStreamReadingError(message string) StreamReadingError {
	return StreamReadingError{message: message}
}

func (p StreamReadingError) Error() string {
	return "stream reading error: " + p.message
}
