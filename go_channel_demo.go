package main

import (
	"fmt"
	"time"
)

func printHello(boolchan chan bool, strchannel chan string) {
	time.Sleep(4 * time.Second)
	status := <-boolchan
	if status {
		strchannel <- "hello"
	}
}
func blockTime(channel chan int) {
	time.Sleep(5 * time.Second)
	channel <- 1
}
func checkStatus(channel chan int, boolchan chan bool) {
	num := <-channel
	if num == 1 {
		boolchan <- true
	}
}
func main() {
	channel := make(chan int)
	boolchan := make(chan bool)
	strchannel := make(chan string)
	go printHello(boolchan, strchannel)
	go blockTime(channel)
	go checkStatus(channel, boolchan)
	fmt.Println(<-strchannel)
}
