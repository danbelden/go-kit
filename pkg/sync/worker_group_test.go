package sync

import (
	"fmt"
	coreCync "sync"
	"testing"
	"time"
)

func TestWorkerGroup_Wait_NoError(t *testing.T) {
	wg := NewWorkerGroup()

	var counter int
	lock := &coreCync.Mutex{}

	f := func() error {
		lock.Lock()
		defer lock.Unlock()

		counter++

		return nil
	}

	for i := 0; i < 10; i++ {
		wg.Add(f)
	}
	wg.Wait()

	if counter != 10 {
		t.Errorf("got: %d, want: 10", counter)
	}
}

func TestWorkerGroup_Wait_WithError(t *testing.T) {
	wg := NewWorkerGroup()

	var counter int
	lock := &coreCync.Mutex{}

	f := func() error {
		lock.Lock()
		defer lock.Unlock()

		counter++
		if counter > 5 {
			return fmt.Errorf("counter is %d", counter)
		}

		return nil
	}

	for i := 0; i < 10; i++ {
		wg.Add(f)
	}
	wg.Wait()

	if counter != 10 {
		t.Errorf("got: %d, want: 10", counter)
	}
}

func TestWorkerGroup_WaitOnError_WithError(t *testing.T) {
	wg := NewWorkerGroup()

	var counter int
	lock := &coreCync.Mutex{}

	f := func() error {
		lock.Lock()
		isSleep := counter > 5
		lock.Unlock()

		if isSleep {
			time.Sleep(time.Second)
		}

		lock.Lock()
		defer lock.Unlock()

		counter++
		if counter > 5 {
			return fmt.Errorf("counter is %d", counter)
		}

		return nil
	}

	for i := 0; i < 10; i++ {
		wg.Add(f)
	}

	if err := wg.WaitOnError(); err == nil {
		t.Fatalf("expected an error")
	}

	lock.Lock()
	defer lock.Unlock()

	if counter == 10 {
		t.Errorf("got: %d, want: <10", counter)
	}
}

func TestWorkerGroup_Add_AfterWait(t *testing.T) {
	wg := NewWorkerGroup()

	var counter int
	lock := &coreCync.Mutex{}

	f := func() error {
		lock.Lock()
		defer lock.Unlock()

		counter++
		time.Sleep(time.Microsecond)

		return nil
	}

	for i := 0; i < 5; i++ {
		wg.Add(f)
	}

	waitChan := make(chan bool)
	go func() {
		wg.Wait()
		waitChan <- true
	}()

	for i := 0; i < 5; i++ {
		wg.Add(f)
	}

	<-waitChan

	if counter != 10 {
		t.Errorf("got: %d, want: 10", counter)
	}
}

func TestWorkerGroup_WaitOnError_WithPanic(t *testing.T) {
	wg := NewWorkerGroup()

	var counter int
	lock := &coreCync.Mutex{}

	f := func() error {
		lock.Lock()
		isSleep := counter > 5
		lock.Unlock()

		if isSleep {
			time.Sleep(time.Second)
		}

		lock.Lock()
		defer lock.Unlock()

		counter++
		if counter > 5 {
			panic(fmt.Sprintf("counter is %d", counter))
		}

		return nil
	}

	for i := 0; i < 10; i++ {
		wg.Add(f)
	}

	err := wg.WaitOnError()
	if err == nil {
		t.Fatalf("expected an error")
	}

	if err.Error() != "worker panic" {
		t.Errorf("unexpected error: %s", err.Error())
	}

	lock.Lock()
	defer lock.Unlock()

	if counter == 10 {
		t.Errorf("got: %d, want: <10", counter)
	}
}

func TestWorkerGroup_GetResults(t *testing.T) {
	wg := NewWorkerGroup()

	var counter int
	lock := &coreCync.Mutex{}

	for i := 0; i < 5; i++ {
		workerID := i + 1
		f := func() error {
			lock.Lock()
			defer lock.Unlock()

			time.Sleep(time.Microsecond)

			counter++

			switch workerID {
			case 3:
				panic(fmt.Sprintf("counter is 3"))
			case 4:
				return fmt.Errorf("counter is 4")
			default:
				return nil
			}
		}
		wg.Add(f)
	}

	wg.Wait()

	results := wg.GetResults()

	numResults := len(results)
	if numResults != 5 {
		t.Fatalf("unexpected result count: %d", numResults)
	}

	for _, result := range results {
		workerID := result.GetWorkerID()
		switch workerID {
		case 1, 2, 5:
			if err := result.GetError(); err != nil {
				t.Errorf("unexpected error for worker %d: %s", workerID, err.Error())
			}
			if p := result.GetPanic(); p != nil {
				t.Errorf("unexpected panic for worker %d", workerID)
			}
			if status := result.GetStatus(); status != WorkerResultStatusDone {
				t.Errorf("unexpected status for worker %d: %s", workerID, status)
			}
		case 3:
			if err := result.GetError(); err == nil {
				t.Errorf("expected an error from worker 3")
			}
			if p := result.GetPanic(); p == nil {
				t.Errorf("expected a panic from worker 3")
			}
			if status := result.GetStatus(); status != WorkerResultStatusPanic {
				t.Errorf("unexpected status for worker 3: %s", status)
			}
		case 4:
			if err := result.GetError(); err == nil {
				t.Errorf("expected an error from worker 4")
			}
			if p := result.GetPanic(); p != nil {
				t.Errorf("unexpected panic for worker 4")
			}
			if status := result.GetStatus(); status != WorkerResultStatusError {
				t.Errorf("unexpected status for worker 4: %s", status)
			}
		default:
			t.Errorf("unexpected worker %d", workerID)
		}
	}
}

func TestWorkerGroup_GetResults_WhileRunning(t *testing.T) {
	wg := NewWorkerGroup()

	var counter int
	lock := &coreCync.Mutex{}

	doneChan := make(chan bool)

	f := func() error {
		lock.Lock()
		defer lock.Unlock()

		counter++
		doneChan <- true

		return nil
	}

	for i := 0; i < 10; i++ {
		wg.Add(f)
	}

	var doneCount int
	for <-doneChan {
		doneCount++
		if doneCount == 6 {
			break
		}
	}

	results := wg.GetResults()

	numResults := len(results)
	if numResults != 10 {
		t.Errorf("unexpected num results; got: %d, want: 10", numResults)
	}

	var numRunning int
	for _, result := range results {
		if result.GetStatus() == WorkerResultStatusRunning {
			numRunning++
		}
	}

	if numRunning < 4 || numRunning > 6 {
		t.Errorf("unexpected running results, got: %d", numRunning)
	}
}
