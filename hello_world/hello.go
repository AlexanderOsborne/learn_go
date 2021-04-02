package main

import "fmt"

// func main() {
// 	fmt.Println("Hello World")
// }

func Hello() string {
	return "Hello, World"
}

func main() {
	fmt.Println(Hello())
}
