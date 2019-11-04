package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second * 2)
	fmt.Printf("Worker %d done\n", id)

	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 30; i++ {
		wg.Add(1)
		time.Sleep(time.Second)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Print("\nDone!\n")
}
