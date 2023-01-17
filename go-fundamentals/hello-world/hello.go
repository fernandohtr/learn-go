package main

import "fmt"

const englishPrefix = "Hello, "
const portuguesePrefix = "Olá, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	if language == "Portuguese" {
		return portuguesePrefix + name
	}
	return englishPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
