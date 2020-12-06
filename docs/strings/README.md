## Strings

### UniqueSlice

```
package main

import "github.com/danbelden/go-kit/pkg/strings"

func main() {
    stringslice := []string{"A", "A", "B", "B", "C", "C"}
    
    uniqueStringSlice := strings.UniqueSlice(slice)
    fmt.Println(uniqueStringSlice)
}
```
```
$ go run main.go
[A, B, C]
```

### SearchWord

```
package main

import "github.com/danbelden/go-kit/pkg/strings"

func main() {
    text := "the quick brown fox"
    word := "quick"

    wordExists := strings.SearchWord(text, word)
    fmt.Println(wordExists)
}
```
```
$ go run main.go
true
```