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
			return nil
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

	"github.com/danbelden/go-kit/pkg/sync"
)

func main() {
	wg := sync.NewWorkerGroup()

	for i := 0; i < 3; i++ {
		wg.Add(func() error {
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
error
```

#### WorkerGroup.GetResults

```
package main

import (
	"fmt"

	"github.com/danbelden/go-kit/pkg/sync"
)

func main() {
	wg := sync.NewWorkerGroup()

	for i := 0; i < 10; i++ {
		workerID := i + 1
		wg.Add(func() error {
			time.Sleep(time.Microsecond)
			switch workerID {
			case 1:
				return fmt.Errorf("error")
			case 2:
				panic("panic")
			default:
				return nil
			}
		})
	}

	time.Sleep(time.Microsecond * 5)
	statusCounts := make(map[string]int64)

	results := wg.GetResults()
	for _, result := range results {
		status := string(result.GetStatus())
		statusCounts[status]++
	}

	fmt.Println(statusCounts)
}
```
```
$ go run main.go
map[done:5 error:1 panic:1 running:3]
```