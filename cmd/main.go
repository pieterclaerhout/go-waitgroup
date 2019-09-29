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
