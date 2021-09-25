## Ints

### UniqueSlice

```
package main

import "github.com/danbelden/go-kit/pkg/ints"

func main() {
	intSlice := []string{1, 1, 2, 2, 3, 3}

	uniqueIntSlice := ints.UniqueSlice(intSlice)
	fmt.Println(uniqueIntSlice)
}
```
```
$ go run main.go
[1, 2, 3]
```