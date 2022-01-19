package main

import "fmt"

const englishHelloPrefix = "Hello"
const frenchHelloPrefix = "Bonjour"
const spanishHelloPrefix = "Hola"

func prefix(language string) string {
	switch language {
	case "English":
		return englishHelloPrefix
	case "French":
		return frenchHelloPrefix
	case "Spanish":
		return spanishHelloPrefix
	default:
		return englishHelloPrefix
	}
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
