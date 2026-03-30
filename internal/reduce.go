package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

func Reduce() {

}

func ReduceWorker(reduceIdx, nMap int, reducef ReduceFunc) {

	grouped := Group(reduceIdx, nMap)

	fileName := fmt.Sprintf("output-%d", reduceIdx)
	f, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer f.Close()

	for key, values := range grouped {
		result := reducef(key, values)
		_, err := fmt.Fprintf(f, "%v %v\n", key, result)
		if err != nil {
			continue
		}
	}

	return
}

func Group(reduceIdx, nMap int) map[string][]string {

	grouped := make(map[string][]string)

	for i := 0; i < nMap; i++ {
		fileName := fmt.Sprintf("mr-%d-%d", i, reduceIdx)

		// Open the file and skip if not available
		f, err := os.Open(fileName)
		if err != nil {
			continue
		}

		// Decode file and add to the grouped
		dec := json.NewDecoder(f)
		for {
			var kv KeyValue
			if err := dec.Decode(&kv); err != nil {
				break
			}
			grouped[kv.Key] = append(grouped[kv.Key], kv.Value)
		}

		f.Close()
	}

	return grouped
}
