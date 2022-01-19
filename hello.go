package main

import "fmt"

const englishHelloPrefix = "Hello"

func Hello(name string) string {
	return fmt.Sprintf("%s, %s", englishHelloPrefix, name)
}
