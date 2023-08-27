package main

import (
	"fmt"
)

func main() {
	quit := make(chan string)
	ch := generator("hello", quit)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	quit <- "bye"
	fmt.Printf("Generator says %s", <-quit)
}

func generator(msg string, quit chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case ch <- fmt.Sprintf("%s %d", msg, i):
			case <-quit:
				quit <- "See you!"
				return
			}
		}
	}()
	return ch
}
