package main

import "fmt"

func Hello(input string) string {
	if input == "" {
		input = "World"
	}
	return "Hello, " + input
}

func main() {
	fmt.Println(Hello("Chris"))
}
