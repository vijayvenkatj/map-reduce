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

func CreateWorker(id int, addr string, nMap, nReduce int) *Worker {

	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	return &Worker{
		ID:      id,
		Client:  client,
		NMap:    nMap,
		NReduce: nReduce,
	}
}

func (w *Worker) Run(mapf MapFunc, reducef ReduceFunc) {
	for {
		// Replace with RPC
		var task = w.getTask()

		switch task.Type {

		case MapTask:
			file := fmt.Sprintf("input-%d.txt", task.ID)
			Map(file, task.ID, w.NReduce, mapf)
			var reply bool
			w.Client.Call("Master.TaskDone", task, &reply)
			break
		case ReduceTask:
			Reduce(task.ID, w.NMap, reducef)
			var reply bool
			w.Client.Call("Master.TaskDone", task, &reply)
			break
		case WaitTask:
			log.Println("No work assigned. Sleeping...(1s)")
			time.Sleep(1 * time.Second)
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
