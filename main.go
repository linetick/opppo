package main

import (
	"fmt"
	"os"
	"pr1/entites/process"
)

func main() {
	if len(os.Args) < 2 {
		//CreateSampleFile()
		os.Exit(1)
	}

	filename := os.Args[1]
	container := process.NewContainer()

	fmt.Printf("Processing file: %s\n", filename)
	if err := process.ReadAndProcessFile(filename, container); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

