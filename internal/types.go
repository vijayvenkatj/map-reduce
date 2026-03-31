package internal

// MasterParams assumes NMap == NumTasks.
type MasterParams struct {
	ID      int
	NMap    int
	NReduce int
}

type KeyValue struct {
	Key   string
	Value string
}

type MapFunc func(string, string) []KeyValue
type ReduceFunc func(string, []string) string

type TaskStatus int

const (
	Idle       TaskStatus = iota
	InProgress            = 1
	Completed             = 2
)

type TaskType int

const (
	WaitTask   TaskType = iota
	MapTask             = 1
	ReduceTask          = 2
	ExitTask            = 3
)

type Task struct {
	ID     int
	Status TaskStatus
	Type   TaskType
}
