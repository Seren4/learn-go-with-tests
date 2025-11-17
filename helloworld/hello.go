package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	italianHelloPrefix = "Buongiorno, "
	spanishLanguage    = "Spanish"
	frenchLanguage     = "French"
	italianLanguage    = "Italian"
)

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanishLanguage:
		prefix = spanishHelloPrefix
	case frenchLanguage:
		prefix = frenchHelloPrefix
	case italianLanguage:
		prefix = italianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
