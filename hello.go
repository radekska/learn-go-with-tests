package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		return fmt.Sprintf("%sWorld!", englishHelloPrefix)
	}
	return fmt.Sprintf("%s%s!", englishHelloPrefix, name)
}

func main() {
	fmt.Println(Hello("Chris"))
}
