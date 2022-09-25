package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func getGreetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%s%s!", getGreetingPrefix(language), name)
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
