## Sync

### WorkerGroup

#### WorkerGroup.Wait

```
package main

import (
	"fmt"

	"github.com/danbelden/go-kit/pkg/sync"
)

func main() {
	wg := sync.NewWorkerGroup()

	for i := 0; i < 3; i++ {
		wg.Add(func() error {
			fmt.Println("test")
			return fmt.Errorf("error")
		})
	}

	wg.Wait()
}
```
```
$ go run main.go
test
test
test
```

#### WorkerGroup.WaitOnError

```
package main

import (
	"fmt"
	"time"

	"github.com/danbelden/go-kit/pkg/sync"
)

func main() {
	wg := sync.NewWorkerGroup()

	for i := 0; i < 3; i++ {
		iNum := i + 1    
		wg.Add(func() error {
			if iNum != 1 {
				time.Sleep(time.Second)
			}
			fmt.Println("test")
			return fmt.Errorf("error")
		})
	}

	if err := wg.WaitOnError(); err != nil {
		fmt.Println(err.Error())
	}
}
```
```
$ go run main.go
test
error
```

#### WorkerGroup.GetResults

```
package main

import (
	"fmt"
	"time"

	"github.com/danbelden/go-kit/pkg/sync"
)

func main() {
	wg := sync.NewWorkerGroup()

	var workerIDs []int64
	for i := 0; i < 5; i++ {
		iNum := i + 1
		workerID := wg.Add(func() error {
			time.Sleep(time.Microsecond * 2)
			switch iNum {
			case 1:
				return fmt.Errorf("error message")
			case 2:
				panic("panic message")
			default:
				return nil
			}
		})
		workerIDs = append(workerIDs, workerID)
		fmt.Println(fmt.Sprintf("worker started: %d", workerID))
	}

	time.Sleep(time.Microsecond)

	results := wg.GetResults()
	for _, workerID := range workerIDs {
		result, ok := results[workerID]
		if !ok {
			continue
		}

		workerID = result.GetWorkerID()
		status := string(result.GetStatus())

		fmt.Println(fmt.Sprintf("worker %d: status > %s", workerID, status))

		workerErr := result.GetError()
		if workerErr != nil {
			fmt.Println(fmt.Sprintf("worker %d: error > %s", workerID, workerErr.Error()))
		}

		workerPanic := result.GetPanic()
		if workerPanic != nil {
			fmt.Println(fmt.Sprintf("worker %d: panic > %s", workerID, workerPanic.(string)))
		}
	}
}
```
```
$ go run main.go
worker started: 1
worker started: 2
worker started: 3
worker started: 4
worker started: 5
worker 1: status > error
worker 1: error > error message
worker 2: status > panic
worker 2: error > worker panic
worker 2: panic > panic message
worker 3: status > done
worker 4: status > done
worker 5: status > running
```