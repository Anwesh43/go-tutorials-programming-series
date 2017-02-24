package main

import "fmt"

func display(name string) (string, string) {
	hello := "hello " + name
	hi := "hi " + name
	return hello, hi
}
func main() {
	hello, hi := display("dog")
	fmt.Println(hello)
	fmt.Println(hi)
}
