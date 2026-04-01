package user

import (
	"strconv"
	"strings"

	"github.com/vijayvenkatj/map-reduce/internal"
)

func MapF(filename string, contents string) []internal.KeyValue {
	kva := []internal.KeyValue{}

	nums := strings.Fields(contents)

	for _, n := range nums {
		kva = append(kva, internal.KeyValue{
			Key:   n,
			Value: "1",
		})
	}

	return kva
}

func ReduceF(key string, values []string) string {
	return strconv.Itoa(len(values))
}
