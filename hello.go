package main

import "fmt"

func main() {
	fmt.Println(Hello("Tri"))
}

func Hello(name string) string {
	return "Hello, " + name
}
