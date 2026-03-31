package internal

type Master struct {
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
	*reply = Task{
		ID:     m.ID,
		Type:   WaitTask,
		Status: Idle,
	}
	return nil
}
