package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const spanishLanguage = "Spanish"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanishLanguage {
		return spanishHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
