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
	// Create Channel
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}
	// for i := 0; i < len(links); i++ {
	// 	// fmt.Println(<-c)
	// 	go checkLink(<-c, c)
	// }

	// Infinite Loop
	// for {
	// 	go checkLink(<-c, c)
	// }

	// Alternative Loop syntax
	// Wait for channel to return some value, once received assign the value to l
	for l := range c {
		// go checkLink(l, c)
		// Function Literal
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// c <- "from channel: Might be down I think"
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	// c <- "from channel: Yes its up"
	c <- link
}
