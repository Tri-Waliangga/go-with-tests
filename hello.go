package main

import "fmt"

func main() {
	fmt.Println(Hello("Tri"))
}

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	return englishHelloPrefix + name
}
