package main

import (
	"strconv"
	"strings"

	"github.com/vijayvenkatj/map-reduce/internal"
)

func main() {

	var mapf internal.MapFunc
	mapf = func(filename string, contents string) []internal.KeyValue {
		var kva []internal.KeyValue

		words := strings.Fields(contents)

		for _, w := range words {
			kva = append(kva, internal.KeyValue{
				Key:   w,
				Value: "1",
			})
		}
		return kva
	}

	var reducef internal.ReduceFunc
	reducef = func(key string, values []string) string {
		return strconv.Itoa(len(values))
	}

	internal.MapWorker("input.txt", 0, 1, mapf)
	internal.ReduceWorker(0, 1, reducef)
}
