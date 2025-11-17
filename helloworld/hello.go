package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const spanishLanguage = "Spanish"
const frenchLanguage = "French"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanishLanguage {
		return spanishHelloPrefix + name
	} else if language == frenchLanguage {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
