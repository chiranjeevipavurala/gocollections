package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Helper function to run benchmarks and save results
func runBenchmarks() {
	fmt.Println("Running Go Collections Benchmarks...")
	fmt.Println("=====================================")

	// Run benchmarks using go test
	cmd := exec.Command("go", "test", "-bench=.", "-benchmem", "./benchmark/")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("\nExecuting: go test -bench=. -benchmem ./benchmark/")
	fmt.Println("--------------------------------------------------")

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running benchmarks: %v\n", err)
		return
	}

	fmt.Println("\nBenchmarking completed!")
}

// Helper function to run specific benchmark categories
func runSpecificBenchmarks(pattern string) {
	fmt.Printf("Running benchmarks matching pattern: %s\n", pattern)
	fmt.Println("================================================")

	cmd := exec.Command("go", "test", "-bench="+pattern, "-benchmem", "./benchmark/")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running benchmarks: %v\n", err)
		return
	}
}

func main() {
	if len(os.Args) > 1 {
		// If arguments provided, run specific benchmarks
		pattern := os.Args[1]
		runSpecificBenchmarks(pattern)
	} else {
		// Run all benchmarks
		runBenchmarks()
	}
}
