package main

import "fmt"

func Hello(arg string) string {
	return "Hello, " + arg
}

func main() {
	fmt.Println(Hello("world"))
}
