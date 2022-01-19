package main

import "fmt"

const englishHelloPrefix = "Hello"
const frenchHelloPrefix = "Bonjour"
const spanishHelloPrefix = "Hola"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "" {
		language = "English"
	}
	prefix := ""
	switch language {
	case "English":
		prefix = englishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return fmt.Sprintf("%s, %s", prefix, name)
}
