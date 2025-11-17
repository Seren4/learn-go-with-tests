package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const italianHelloPrefix = "Buongiorno, "
const spanishLanguage = "Spanish"
const frenchLanguage = "French"
const italianLanguage = "Italian"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	switch language {

	case spanishLanguage:
		return spanishHelloPrefix + name
	case frenchLanguage:
		return frenchHelloPrefix + name
	case italianLanguage:
		return italianHelloPrefix + name
	default:
		return englishHelloPrefix + name
	}
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
