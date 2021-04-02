package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHolaPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return spanishHolaPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
