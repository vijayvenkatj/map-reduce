package internal

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"log"
	"math"
	"os"
)

func iHash(str string) int {
	hash := fnv.New32()
	_, err := hash.Write([]byte(str))
	if err != nil {
		log.Printf("Error hashing the string %s", str)
		return math.MaxInt
	}
	return int(hash.Sum32() & 0x7FFFFF)
}

/*
Map function, Takes in filenames, number of Map workers and the user mapf as arguments.
*/

func Map() {
	// Take file names
}

/*
MapWorker takes a file as input and applies user defined mapf to write intermediate file based on partitioning function.
*/
func MapWorker(fileName string, mapIdx, nReduce int, mapf MapFunc) {
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	kva := mapf(fileName, string(fileData))

	files := make([]*os.File, nReduce)
	encoders := make([]*json.Encoder, nReduce)

	for i := 0; i < nReduce; i++ {
		tmp := fmt.Sprintf("mr-%d-%d.tmp", mapIdx, i)

		f, err := os.Create(tmp)
		if err != nil {
			panic(err)
		}

		files[i] = f
		encoders[i] = json.NewEncoder(f)
	}

	for _, kv := range kva {
		r := iHash(kv.Key) % nReduce
		encoders[r].Encode(&kv)
	}

	for i := 0; i < nReduce; i++ {
		files[i].Close()

		final := fmt.Sprintf("mr-%d-%d", mapIdx, i)
		tmp := fmt.Sprintf("mr-%d-%d.tmp", mapIdx, i)

		os.Rename(tmp, final)
	}
}
