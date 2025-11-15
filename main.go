package main

import (
	"fmt"
	_ "net/http/pprof"
	"os"
	"pr1/entites/container"
	"pr1/entites/process"
	"runtime/pprof"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: pr1 <input_file>")
		os.Exit(1)
	}

	c := container.New()
	if err := process.ExecuteFile(os.Args[1], c); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Total figures processed: %d\n", c.Count())

	allocsFile, err := os.Create("allocs.pprof")
	if err != nil {
		fmt.Printf("Failed to create allocs profile: %v\n", err)
		os.Exit(1)
	}
	defer allocsFile.Close()

	pprof.Lookup("allocs").WriteTo(allocsFile, 0)
	fmt.Println("Allocs profile saved")

	//runtime.KeepAlive(container)
}