package main

import "fmt"

const prefix = "Hello, "

func Hello(arg string) string {
	if arg == "" {
		arg = "world"
	}

	return prefix + arg
}

func main() {
	fmt.Println(Hello("world"))
}
