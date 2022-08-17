package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

/*
	We can WaitGroups to clean up the example a bit
*/

func main() {
	wg := new(sync.WaitGroup)

	wg.Add(2)

	go fastProcess(wg)
	go slowProcess(wg)

	wg.Wait()

	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("Both processes have finished!")
	fmt.Println(strings.Repeat("-", 50))
}

func fastProcess(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 4; i++ {
		time.Sleep(time.Second)
		fmt.Println("Fast counter: ", i)
	}
}

func slowProcess(wg *sync.WaitGroup) {
	defer wg.Done()

	for j := 0; j < 4; j++ {
		time.Sleep(2 * time.Second)
		fmt.Println("Slow counter: ", j)
	}
}
