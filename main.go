package main

import (
	"fmt"
	_ "net/http/pprof"
	"os"
	"pr1/entites/process"
	"runtime"
	"runtime/pprof"
)

func main() {

	if len(os.Args) < 2 {
		os.Exit(1)
	}
	filename := os.Args[1]
	container := process.NewContainer()

	fmt.Printf("Processing file: %s\n", filename)
	if err := process.ReadAndProcessFile(filename, container); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Total figures processed: %d\n", container.Count())

	allocsFile, err := os.Create("allocs.pprof")
	if err != nil {
		fmt.Printf("Failed to create allocs profile: %v\n", err)
		os.Exit(1)
	}
	defer allocsFile.Close()

	pprof.Lookup("allocs").WriteTo(allocsFile, 0)
	fmt.Println("Allocs profile saved")

	runtime.KeepAlive(container)
}