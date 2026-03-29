package internal

import (
	"os"
	"strings"
)

func MapReduce(inputFile string, nMap, nReduce int, mapf MapFunc, reducef ReduceFunc) string {

	fileData, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	content := string(fileData)
	data := strings.Split(content, "\n")

	n := len(data)
	chunkSize := (n + nMap - 1) / nMap

	var chunks [][]string

	for i := 0; i < n; i += chunkSize {
		end := min(i+chunkSize, n)
		chunks = append(chunks, data[i:end])
	}

	// Use Map function

	// Group into maps

	// Use Reduce function

	// Return the result

}
