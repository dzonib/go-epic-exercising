package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	userName := fetchUser()

	respChan := make(chan any, 2)

	wg := &sync.WaitGroup{}

	wg.Add(2)

	go fetchUserLikes(userName, respChan, wg)
	go fetchUserMatch(userName, respChan, wg)

	wg.Wait() // block until 2 wg.Done()

	close(respChan)

	for resp := range respChan {
		// check for type if its any type
		// likes, ok := resp.(int)

		// if ok {
		// 	 it is an int
		// }

		fmt.Println("resp:", resp)
	}

	// fmt.Println("Likes:", respChan)
	// fmt.Println("Match:", respChan)

	fmt.Println("Took us", time.Since(start))
}

func fetchUser() string {

	time.Sleep(time.Millisecond * 100)
	return "BOB"
}

func fetchUserLikes(userName string, respChan chan any, wg *sync.WaitGroup) {

	time.Sleep(time.Millisecond * 150)
	respChan <- 11
	wg.Done()
}

func fetchUserMatch(userName string, respChan chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	respChan <- "Anna"
	wg.Done()
}
