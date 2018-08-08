# go-waitgroup

[![GoDoc](https://godoc.org/github.com/pieterclaerhout/go-waitgroup?status.png)](https://godoc.org/github.com/pieterclaerhout/go-waitgroup)

An package that allows you to use the constructs of a [`sync.WaitGroup`](https://golang.org/pkg/sync/#WaitGroup) to
create a pool of goroutines and control the concurrency.

Using it is just like a normal [`sync.WaitGroup`](https://golang.org/pkg/sync/#WaitGroup). The only difference is the initialisation. When you use `waitgroup.NewWaitGroup`, you have the option to specify it's size.

Any `int` which is bigger than `0` will limit the number of concurrent goroutines. If you specify `-1` or `0`, all goroutines will run at once (just like a plain [`sync.WaitGroup`](https://golang.org/pkg/sync/#WaitGroup)).

```go
package main

import (
    "fmt"
    "github.com/pieterclaerhout/go-waitgroup"
)

func main() {
    
    urls := []string{
        "https://www.easyjet.com/",
        "https://www.skyscanner.de/",
        "https://www.ryanair.com",
        "https://wizzair.com/",
        "https://www.swiss.com/",
    }

    wg := waitgroup.NewWaitGroup(3)

    for _, url := range urls {
        wg.BlockAdd()
        go func(url string) {
            fmt.Println("%s: checking", url)
            _, err := http.Get(url)
            fmt.Println("%s: result: %v", err)
            wg.Done()
        }(url)
    }

    wg.Wait()
    fmt.Println("Finished")

}
```
