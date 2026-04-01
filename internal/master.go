package internal

import (
	"sync"
)

type Master struct {
	mu sync.Mutex

	MapTasks    []Task
	ReduceTasks []Task

	ID      int
	Phase   string
	NMap    int
	NReduce int
}

func CreateMaster(params MasterParams) *Master {

	mapTask := make([]Task, params.NMap)
	reduceTask := make([]Task, params.NReduce)

	for i := 0; i < params.NMap; i++ {
		mapTask[i] = Task{
			ID:     i,
			Type:   MapTask,
			Status: Idle,
		}
	}
	for i := 0; i < params.NReduce; i++ {
		reduceTask[i] = Task{
			ID:     i,
			Type:   ReduceTask,
			Status: Idle,
		}
	}

	return &Master{
		MapTasks:    mapTask,
		ReduceTasks: reduceTask,

		ID:      params.ID,
		Phase:   "map",
		NMap:    params.NMap,
		NReduce: params.NReduce,
	}
}

func (m *Master) Task(args *int, reply *Task) error {

	m.mu.Lock()
	defer m.mu.Unlock()

	// Check for the phase
	if m.Phase == "map" {
		for i := 0; i < m.NMap; i++ {
			if m.MapTasks[i].Status == Idle {
				m.MapTasks[i].Status = InProgress
				*reply = m.MapTasks[i]
				return nil
			}
		}
		*reply = Task{
			ID:     -1,
			Type:   WaitTask,
			Status: Idle,
		}
		return nil
	}

	if m.Phase == "reduce" {
		for i := 0; i < m.NReduce; i++ {
			if m.ReduceTasks[i].Status == Idle {
				m.ReduceTasks[i].Status = InProgress
				*reply = m.ReduceTasks[i]
				return nil
			}
		}
		*reply = Task{
			ID:     -1,
			Type:   WaitTask,
			Status: Idle,
		}
		return nil
	}

	if m.Phase == "completed" {
		*reply = Task{
			ID:     -2,
			Type:   ExitTask,
			Status: Idle,
		}
		return nil
	}

	return nil
}

func (m *Master) TaskDone(args *Task, reply *bool) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if args.Type == MapTask {
		m.MapTasks[args.ID].Status = Completed

		for i := 0; i < m.NMap; i++ {
			if m.MapTasks[i].Status != Completed {
				*reply = true
				return nil
			}
		}
		m.Phase = "reduce"
	}

	if args.Type == ReduceTask {
		m.ReduceTasks[args.ID].Status = Completed
		for i := 0; i < m.NReduce; i++ {
			if m.ReduceTasks[i].Status != Completed {
				*reply = true
				return nil
			}
		}
		m.Phase = "completed"
	}

	*reply = true
	return nil
}
