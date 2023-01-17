package main

import "fmt"

const englishPrefix = "Hello, "
const portuguesePrefix = "Olá, "
const spanishPrefix = "Hola, "

func Hello(name string, language string) string {
	var prefix string

	if name == "" {
		name = "world"
	}

	switch language {
	case "Portuguese":
		prefix = portuguesePrefix
	case "Spanish":
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
