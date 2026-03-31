package internal

import (
	"log"
	"time"
)

type Worker struct {
	ID int
}

func CreateWorker(id int) *Worker {
	return &Worker{
		ID: id,
	}
}

func (w *Worker) Run() {
	for {
		// Replace with RPC
		var rpcResult Task

		switch rpcResult.Type {

		case MapTask:
			break
		case ReduceTask:
			break
		case WaitTask:
			log.Println("No work assigned. Sleeping...(2s)")
			time.Sleep(2 * time.Second)
			break
		case ExitTask:
			log.Println("Worker has been shut down!")
			return
		}

	}
}
