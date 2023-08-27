package main

// concurrency code

import (
	"fmt"
	"sync"
	"time"
)

func downloadJSON(u string) {
	println(u)
	time.Sleep(time.Second * 6)
}

func main() {
	before := time.Now()

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()

			u := fmt.Sprintf("http://www.example.com/%d", i)
			downloadJSON(u)
		}()
	}
	wg.Wait()

	fmt.Printf("%v\n", time.Since(before))
}
