package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	signal := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go myFunc(signal, &wg)
	go myfunc2(signal)
	go myfunc3(signal)
	// signal <- true
	signal <- "first"
	wg.Wait()

}

func myFunc(signal chan string, wg *sync.WaitGroup) {
	count := 0
	for {
		if "first" == <-signal {
			fmt.Print("first\t")
		}
		signal <- "second"
		time.Sleep(time.Second * 2)

		if count++; count == 10 {
			wg.Done()
			return
		}
	}
}

func myfunc2(signal chan string) {
	for {
		if "second" == <-signal {
			fmt.Print("Second\t")
		}
		signal <- "third"
		time.Sleep(time.Second)
	}
}

func myfunc3(signal chan string) {
	for {
		if "third" == <-signal {
			fmt.Print("third\t")
		}
		signal <- "first"
		time.Sleep(time.Second)
	}
}
