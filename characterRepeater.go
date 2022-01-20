package main

import "fmt"

func RepeatCharacter(c byte, numberOfRepetitions int) string {
	result := ""
	for i := 0; i < numberOfRepetitions; i++ {
		result += string(c)
	}
	return fmt.Sprintf("%s", result)
}
