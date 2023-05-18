package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Channels are used to communicate between routines
	c := make(chan string)

	for _, link := range links {
		// 'go' is keyword for run code in a new routine (Child Routine)
		// Check Go Scheduler -> 1 CPU = 1 Routine rouning
		// Parallelism -> Multiple things at the same time
		// Concurrency -> Not necessarily
		go checkLink(link, c)
	}

	// Check one time for each link and done
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// Receiving channel values - e.g: value <- channel
	for l := range c {
		// Anonymous function (Javascript) - Lambda (Python)
		// Try to avoid referencing the same var in 2 routines - Use Channels and Parameters
		go func(link string) {
			// Check continuously
			// Small delay
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) // Using funtion
	}
}

func checkLink(link string, c chan string) {
	// Get is waiting for response normally - Need concurrency to avoid waiting !
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// Send data to channel
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	// Send data to channel
	c <- link
}
