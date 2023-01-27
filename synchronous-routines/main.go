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


// using Waitgroups only

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// type myWaitgrp struct {
// 	counter int
// 	wg      *sync.WaitGroup
// }

// // Date should be in the Format MM-DD-YYYY
// // for single digits add 0 in front of the Month or Date, for example MM -> 02 or DD-> 04 and Year -> 0202
// // for year `202`.

// func main() {
// 	var wg1 sync.WaitGroup
// 	var wg2 sync.WaitGroup
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	wg2.Add(1)
// 	go myfunc1(&wg1, &wg2)
// 	go myfunc2(&wg1, &wg2, &wg)
// 	wg.Wait()
// }

// func myfunc2(waitGroup1 *sync.WaitGroup, waitGroup2, wg *sync.WaitGroup) {
// 	t := time.Now()
// 	for {
// 		waitGroup1.Wait()
// 		fmt.Print("1st go-rouitne \t")
// 		waitGroup1.Add(1)
// 		waitGroup2.Done()
// 		time.Sleep(time.Second * 2)
// 		since := time.Since(t).Seconds()
// 		if time.Duration(since) > 10 {
// 			wg.Done()
// 			break
// 		}
// 	}
// }

// func myfunc1(waitGroup1 *sync.WaitGroup, waitGroup2 *sync.WaitGroup) {
// 	for {
// 		waitGroup2.Wait()
// 		fmt.Print("2nd go-rouitne \t")
// 		waitGroup2.Add(1)
// 		waitGroup1.Done()
// 	}
// }
