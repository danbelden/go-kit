package sync

type WorkerStatus string

const (
	WorkerResultStatusRunning WorkerStatus = "running"
	WorkerResultStatusError   WorkerStatus = "error"
	WorkerResultStatusPanic   WorkerStatus = "panic"
	WorkerResultStatusDone    WorkerStatus = "done"
)

type WorkerResult interface {
	GetStatus() WorkerStatus
	GetError() error
	GetPanic() interface{}
	GetWorkerID() int64
}

type workerResult struct {
	err      error
	panic    interface{}
	status   WorkerStatus
	workerID int64
}

// GetError returns the error
func (wr *workerResult) GetError() error {
	return wr.err
}

// GetPanic returns the panic
func (wr *workerResult) GetPanic() interface{} {
	return wr.panic
}

// GetStatus returns the status
func (wr *workerResult) GetStatus() WorkerStatus {
	return wr.status
}

// GetWorkerID returns the worker id
func (wr *workerResult) GetWorkerID() int64 {
	return wr.workerID
}
