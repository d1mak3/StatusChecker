package parsers

type NoProductListNodeParsingError struct{}

func (e NoProductListNodeParsingError) Error() string {
	return getError("no product list node")
}

type NoStatusParsingError struct{}

func (e NoStatusParsingError) Error() string {
	return getError("no status")
}

type HtmlParsingError struct{ message string }

func (e HtmlParsingError) Error() string {
	return getError(e.message)
}

func getError(message string) string {
	return "html parsing error: " + message
}
