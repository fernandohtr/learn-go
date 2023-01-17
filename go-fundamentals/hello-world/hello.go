package main

import "fmt"

const englishPrefix = "Hello, "
const portuguesePrefix = "Olá, "
const spanishPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	prefix := englishPrefix

	switch language {
	case "Portuguese":
		prefix = portuguesePrefix
	case "Spanish":
		prefix = spanishPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
