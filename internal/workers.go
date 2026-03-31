package internal

import (
	"log"
	"net/rpc"
	"time"
)

type Worker struct {
	ID     int
	Client *rpc.Client
}

func CreateWorker(id int) *Worker {

	client, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	return &Worker{
		ID:     id,
		Client: client,
	}
}

func (w *Worker) Run() {
	for {
		// Replace with RPC
		var rpcResult = w.getTask()
		
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

func (w *Worker) getTask() Task {
	var task Task
	err := w.Client.Call("Master.Task", w.ID, &task)
	if err != nil {
		log.Println("Call master task failed:", err)
		return Task{}
	}
	return task
}
