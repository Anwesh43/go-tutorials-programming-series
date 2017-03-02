package main

import "fmt"

func printLength(persons map[string]int) {
	fmt.Println("we have", len(persons), "persons")
}
func main() {
	persons := make(map[string]int)
	persons["John"] = 26
	persons["Sam"] = 27
	persons["Arsene"] = 68
	printLength(persons)
	delete(persons, "Arsene")
	printLength(persons)
	for key, value := range persons {
		fmt.Println("Hi person name is", key, "and his/her age is ", value, "years")
	}

}
