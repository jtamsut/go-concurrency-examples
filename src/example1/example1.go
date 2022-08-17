package main

import (
	"fmt"
	"strings"
	"time"
)

/*
	Let's say we have two processes: fastProcess and slowProcess and we want to run them
	each in their own goroutine and be notified when they finish
*/

func main() {
	fastChannel := make(chan int)
	slowChannel := make(chan int)

	defer close(fastChannel)
	defer close(slowChannel)

	go fastProcess(fastChannel)
	go slowProcess(slowChannel)

	select {
	case <-slowChannel:
		fmt.Println("Slow counter has finished!")
	}

	select {
	case <-fastChannel:
		fmt.Println("Slow counter has finished!")
	}

	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("Both processes have finished!")
	fmt.Println(strings.Repeat("-", 50))

}

func fastProcess(ch chan int) {
	for i := 0; i < 4; i++ {
		time.Sleep(time.Second)
		fmt.Println("Fast counter: ", i)
	}

	ch <- 1
}

func slowProcess(ch chan int) {
	for j := 0; j < 4; j++ {
		time.Sleep(2 * time.Second)
		fmt.Println("Slow counter: ", j)
	}

	ch <- 0
}
