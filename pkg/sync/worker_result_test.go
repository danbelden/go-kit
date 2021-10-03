package sync

import (
	"fmt"
	"testing"
)

func TestWorkerResult_GetError(t *testing.T) {
	wr := workerResult{
		err: fmt.Errorf("test"),
	}
	wrErr := wr.GetError()
	if wrErr != wr.err {
		t.Errorf("unexpected error: %s", wrErr)
	}
}

func TestWorkerResult_GetPanic(t *testing.T) {
	wr := workerResult{
		panic: "help",
	}
	wrPanic := wr.GetPanic()
	if wrPanic != wr.panic {
		t.Errorf("unexpected panic: %s", wrPanic)
	}
}

func TestWorkerResult_GetStatus(t *testing.T) {
	wr := workerResult{
		status: WorkerResultStatusDone,
	}
	wrStatus := wr.GetStatus()
	if wrStatus != wr.status {
		t.Errorf("unexpected status: %s", wrStatus)
	}
}

func TestWorkerResult_GetWorkerID(t *testing.T) {
	wr := workerResult{
		workerID: 1,
	}
	wrWorkerID := wr.GetWorkerID()
	if wrWorkerID != wr.workerID {
		t.Errorf("unexpected workerID: %d", wrWorkerID)
	}
}
