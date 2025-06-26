package testing

import (
	"os"
	"parsers"
	"testing"
)

func Test_ParseBrandLinks_ShouldBeOk(t *testing.T) {
	// Arrange
	const expectedLinksCount = 9
	file, _ := os.Open(shopScriptHtmlFile)

	// Act
	links, _ := parsers.GetHtmlResponseParser().ParseBrandLinks(file)

	// Assert
	if len(links) != expectedLinksCount {
		t.Errorf("expected %d link(-s), found %d", expectedLinksCount, len(links))
	}
}

func Test_ParseLinksFromEmptyFile_ShouldBeOk(t *testing.T) {
	// Arrange
	const expectedLinksCount = 0
	file, _ := os.Open(emptyFile)

	// Act
	links, _ := parsers.GetHtmlResponseParser().ParseBrandLinks(file)

	// Assert
	if len(links) != 0 {
		t.Errorf("expected %d link(-s), found %d", expectedLinksCount, len(links))
	}
}

func Test_ParseProductLinks_ShouldBeOk(t *testing.T) {
	// Arrange
	const expectedLinksCount = 30
	file, _ := os.Open(pitakaHtmlFile)

	// Act
	links, _ := parsers.GetHtmlResponseParser().ParseProductLinks(file)

	// Assert
	if len(links) != expectedLinksCount {
		t.Errorf("expected %d link(-s), found %d", expectedLinksCount, len(links))
	}
}

func Test_ParseProductLinks_NoProducts_ShouldReturnErr(t *testing.T) {
	// Arrange
	file, _ := os.Open(shopScriptHtmlFile)
	expectedError := parsers.NoProductListNodeParsingError{}.Error()

	// Act
	_, err := parsers.GetHtmlResponseParser().ParseProductLinks(file)

	// Assert

	if err == nil || err.Error() != expectedError {
		t.Errorf("expected %s error, found %v", expectedError, err)
	}
}

func Test_ParseProductLinks_GarbageLinks_ShouldBeEmpty(t *testing.T) {
	// Arrange
	const expectedLinksCount = 0
	file, _ := os.Open(drugieHtmlFile)

	// Act
	links, _ := parsers.GetHtmlResponseParser().ParseProductLinks(file)

	// Assert
	if len(links) != expectedLinksCount {
		t.Errorf("expected %d link(-s), found %d", expectedLinksCount, len(links))
	}
}

func Test_ParseProductStatus_ShouldBeOk(t *testing.T) {
	// Arrange
	expectedStatuses := []struct {
		status   int
		filePath string
	}{
		{parsers.LowStatus, chekholLowFile},
		{parsers.HighStatus, chekholHighFile},
		{parsers.CriticalStatus, chekholCritFile},
		{parsers.OutOfStockStatus, chekholTranspFile},
	}

	for _, expected := range expectedStatuses {
		file, _ := os.Open(expected.filePath)

		// Act
		status, _ := parsers.GetHtmlResponseParser().ParseProductStatus(file)

		// Assert
		if status != expected.status {
			t.Errorf("expected '%d' status, found '%d'", expected.status, status)
		}
	}
}

func Test_ParseProductStatus_NoStatus_ShouldBeErr(t *testing.T) {
	// Arrange
	expectedError := parsers.NoStatusParsingError{}.Error()
	file, _ := os.Open(shopScriptHtmlFile)

	// Act
	_, err := parsers.GetHtmlResponseParser().ParseProductStatus(file)

	// Assert
	if err == nil || err.Error() != expectedError {
		t.Errorf("expected '%s' status, found '%s'", expectedError, err)
	}
}
