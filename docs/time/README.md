## Time

### Day

```
package main

import (
    coreTime "time"
    "github.com/danbelden/go-kit/pkg/time"
)

func main() {
    timeNow := coreTime.Now()
    timeDay := timeNow.Add(time.Day)
    
    fmt.Println(timeNow)
    fmt.Println(timeDay)
}
```
```
$ go run main.go
2021-10-17 10:12:32.5406 +0100 BST m=+0.000078623
2021-10-18 10:12:32.5406 +0100 BST m=+86400.000078623
```

### Week

```
package main

import (
    coreTime "time"
    "github.com/danbelden/go-kit/pkg/time"
)

func main() {
    timeNow := coreTime.Now()
    timeWeek := timeNow.Add(time.Week)
    
    fmt.Println(timeNow)
    fmt.Println(timeWeek)
}
```
```
$ go run main.go
2021-10-17 10:13:34.218832 +0100 BST m=+0.000074231
2021-10-24 10:13:34.218832 +0100 BST m=+604800.000074231
```