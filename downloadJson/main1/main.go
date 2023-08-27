package main

// original code

import (
	"fmt"
	"time"
)

func downloadJSON(u string) {
	println(u)
	time.Sleep(time.Second * 6)
}

func main() {
	before := time.Now()
	for i := 0; i < 100; i++ {
		u := fmt.Sprintf("http://www.example.com/%d", i)
		downloadJSON(u)
	}

	fmt.Printf("%v\n", time.Since(before))
}
