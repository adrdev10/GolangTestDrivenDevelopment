package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello "

func Hello(name string) string {
	if n := name; n != "" {
		return englishHelloPrefix + n
	}
	return englishHelloPrefix + "World"
}

func main() {
	fmt.Println(Hello("World"))
}
