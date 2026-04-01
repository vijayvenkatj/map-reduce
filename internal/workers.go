package internal

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Worker struct {
	ID     int
	Client *rpc.Client

	NMap    int
	NReduce int
}

func CreateWorker(id int, addr string) *Worker {

	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	return &Worker{
		ID:     id,
		Client: client,
	}
}

func (w *Worker) Run(mapf MapFunc, reducef ReduceFunc) {
	for {
		// Replace with RPC
		var task = w.getTask()

		switch task.Type {

		case MapTask:
			file := fmt.Sprintf("input-%d.txt", w.ID)
			Map(file, task.ID, w.NReduce, mapf)
			break
		case ReduceTask:
			Reduce(task.ID, w.NMap, reducef)
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
