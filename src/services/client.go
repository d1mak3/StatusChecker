package services

import (
	"io"
	"os"
)

func GetClient(allBrandsHtmlFilePath, brandHtmlFilePath, productHtmlFilePath string) Client {
	return fileClient{
		allBrandsHtmlFilePath: allBrandsHtmlFilePath,
		brandHtmlFilePath:     brandHtmlFilePath,
		productHtmlFilePath:   productHtmlFilePath,
	}
}

type Client interface {
	GetAllBrandsReader() (io.Reader, error)
	GetBrandReader(brandName string) (io.Reader, error)
	GetProductReader(productName string) (io.Reader, error)
}

// Debug client
type fileClient struct {
	allBrandsHtmlFilePath, brandHtmlFilePath, productHtmlFilePath string
}

func (client fileClient) GetAllBrandsReader() (io.Reader, error) {
	return os.Open(client.allBrandsHtmlFilePath)
}

func (client fileClient) GetBrandReader(brandName string) (io.Reader, error) {
	return os.Open(client.brandHtmlFilePath)
}

func (client fileClient) GetProductReader(productName string) (io.Reader, error) {
	return os.Open(client.productHtmlFilePath)
}
