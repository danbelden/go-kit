package sync

import (
	"fmt"
	syncCore "sync"
)

type WorkerGroup interface {
	Add(f func() error) int64
	GetResults() map[int64]WorkerResult
	Wait()
	WaitOnError() error
}

type workerGroup struct {
	counter   int64
	doneChans map[int64]chan bool
	errChan   chan error
	lock      *syncCore.RWMutex
	results   map[int64]*workerResult
	waitLock  *syncCore.Mutex
}

// NewWorkerGroup returns a worker group
func NewWorkerGroup() WorkerGroup {
	return &workerGroup{
		doneChans: make(map[int64]chan bool),
		errChan:   make(chan error),
		lock:      &syncCore.RWMutex{},
		results:   make(map[int64]*workerResult),
		waitLock:  &syncCore.Mutex{},
	}
}

// Add registers and launches a given function in the wait group
func (wg *workerGroup) Add(worker func() error) int64 {
	wg.lock.Lock()
	defer wg.lock.Unlock()

	wg.counter++
	workerID := wg.counter
	doneChan := make(chan bool)

	wg.doneChans[workerID] = doneChan
	wg.results[workerID] = &workerResult{
		workerID: workerID,
		status:   WorkerResultStatusRunning,
	}

	wg.runWorker(workerID, worker, wg.results, doneChan, wg.errChan)

	return workerID
}

func (wg *workerGroup) runWorker(
	workerID int64,
	worker func() error,
	results map[int64]*workerResult,
	doneChan chan bool,
	errChan chan error,
) {
	go func() {
		var workerErr error
		workerStatus := WorkerResultStatusDone
		defer func() {
			var workerPanic interface{}
			if r := recover(); r != nil {
				workerErr = fmt.Errorf("worker panic")
				workerPanic = r
				workerStatus = WorkerResultStatusPanic
			}
			wg.lock.Lock()
			if result, ok := results[workerID]; ok {
				result.err = workerErr
				result.panic = workerPanic
				result.status = workerStatus
			}
			wg.lock.Unlock()
			if workerErr != nil {
				errChan <- workerErr
			}
			doneChan <- true
			close(doneChan)
		}()
		if workerErr = worker(); workerErr != nil {
			workerStatus = WorkerResultStatusError
		}
	}()
}

// GetResults returns the map of worker results indexed by workerID
func (wg *workerGroup) GetResults() map[int64]WorkerResult {
	wg.lock.RLock()
	defer wg.lock.RUnlock()

	results := make(map[int64]WorkerResult)
	for key, result := range wg.results {
		results[key] = result
	}

	return results
}

// Wait will block until all workers have finished executing
func (wg *workerGroup) Wait() {
	wg.waitLock.Lock()
	defer wg.waitLock.Unlock()

	_ = wg.processDoneChans(false)
}

// WaitOnError will check, hold and return when all registered functions registered in the group are complete
func (wg *workerGroup) WaitOnError() error {
	wg.waitLock.Lock()
	defer wg.waitLock.Unlock()

	return wg.processDoneChans(true)
}

func (wg *workerGroup) processDoneChans(stopOnError bool) error {
	for {
		workerID, resultChan := wg.getDoneChanToProcess()
		if resultChan == nil {
			return nil
		}
		select {
		case err := <-wg.errChan:
			if stopOnError {
				return err
			}
		case <-resultChan:
			wg.lock.Lock()
			delete(wg.doneChans, workerID)
			wg.lock.Unlock()
		}
	}
}

func (wg *workerGroup) getDoneChanToProcess() (int64, chan bool) {
	wg.lock.RLock()
	defer wg.lock.RUnlock()

	for workerID, doneChan := range wg.doneChans {
		return workerID, doneChan
	}

	return 0, nil
}
