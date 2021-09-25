## Int64s

### UniqueSlice

```
package main

import "github.com/danbelden/go-kit/pkg/int64s"

func main() {
	intSlice := []int64{1, 1, 2, 2, 3, 3}

	uniqueIntSlice := int64s.UniqueSlice(intSlice)
	fmt.Println(uniqueIntSlice)
}
```
```
$ go run main.go
[1, 2, 3]
```