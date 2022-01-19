package main

import "fmt"

const englishHelloPrefix = "Hello"
const frenchHelloPrefix = "Bonjour"
const spanishHelloPrefix = "Hola"

func prefix(language string) (result string) {
	switch language {
	case "English":
		result = englishHelloPrefix
	case "French":
		result = frenchHelloPrefix
	case "Spanish":
		result = spanishHelloPrefix
	default:
		result = englishHelloPrefix
	}
	return
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "" {
		language = "English"
	}
	return fmt.Sprintf("%s, %s", prefix(language), name)
}
