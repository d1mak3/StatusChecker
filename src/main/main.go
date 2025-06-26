package main

import (
	"io"
	"os"
	"services"
)

func main() {
	args := os.Args
	if len(args) != 4 {
		panic("Invalid number of arguments")
	}

	client := services.GetClient(args[1], args[2], args[3])
	allBrandsReader, err := client.GetAllBrandsReader()
	if err != nil {
		panic(err.Error())
	}
}

func getLinks(reader io.Reader) {

}
