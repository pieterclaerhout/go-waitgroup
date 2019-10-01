# go-waitgroup

[![Go Report Card](https://goreportcard.com/badge/github.com/pieterclaerhout/go-waitgroup)](https://goreportcard.com/report/github.com/pieterclaerhout/go-waitgroup)
[![Documentation](https://godoc.org/github.com/pieterclaerhout/go-waitgroup?status.svg)](http://godoc.org/github.com/pieterclaerhout/go-waitgroup)
[![license](https://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://github.com/pieterclaerhout/go-waitgroup/raw/master/LICENSE)
[![GitHub version](https://badge.fury.io/gh/pieterclaerhout%2Fgo-waitgroup.svg)](https://badge.fury.io/gh/pieterclaerhout%2Fgo-waitgroup)
[![GitHub issues](https://img.shields.io/github/issues/pieterclaerhout/go-waitgroup.svg)](https://github.com/pieterclaerhout/go-waitgroup/issues)

An package that allows you to use the constructs of a [`sync.WaitGroup`](https://golang.org/pkg/sync/#WaitGroup) to
create a pool of goroutines and control the concurrency.

Using it is just like a normal [`sync.WaitGroup`](https://golang.org/pkg/sync/#WaitGroup). The only difference is the initialisation. When you use `waitgroup.NewWaitGroup`, you have the option to specify it's size.

Any `int` which is bigger than `0` will limit the number of concurrent goroutines. If you specify `-1` or `0`, all goroutines will run at once (just like a plain [`sync.WaitGroup`](https://golang.org/pkg/sync/#WaitGroup)).

```go
package main

import (
    "fmt"
    "net/http"

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
			defer wg.Done()
			fmt.Printf("%s: checking\n", url)
			res, err := http.Get(url)
			if err != nil {
				fmt.Println("Error: %v")
			} else {
				defer res.Body.Close()
				fmt.Printf("%s: result: %v\n", url, err)
			}
		}(url)
	}

    wg.Wait()
    fmt.Println("Finished")

}
```

There is also a way to use function closures to make it even more readable:

```go
package main

import (
	"fmt"
	"net/http"

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

		urlToCheck := url
		wg.Add(func() {
			fmt.Printf("%s: checking\n", urlToCheck)
			res, err := http.Get(urlToCheck)
			if err != nil {
				fmt.Println("Error: %v")
			} else {
				defer res.Body.Close()
				fmt.Printf("%s: result: %v\n", urlToCheck, err)
			}
		})

	}

	wg.Wait()
	fmt.Println("Finished")

}
```