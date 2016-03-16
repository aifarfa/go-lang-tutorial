package main

import (
	"fmt"
	"time"
)

func main() {
	FeedTextChanel()
	FeedBufferChanel()

	ch := make(chan int, 2)
	go Tricker(ch, 2*time.Second)
	for val := range ch {
		fmt.Println("retrieve value:", val)
	}
}

func FeedTextChanel() {
	// non buffer chanel
	ch := make(chan string)
	go echo(ch) // chaneling with new routine
	ch <- "hello!"
	ch <- "this is a message"
	ch <- "from other routine"
	close(ch)
}

func echo(ch chan string) {
	for msg := range ch {
		fmt.Println(msg)
	}
}

func Tricker(ch chan int, delay time.Duration) {
	for counter := 5; counter > 0; counter-- {
		time.Sleep(delay)
		ch <- counter
	}
	close(ch)
}
