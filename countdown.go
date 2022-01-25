package main

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

func Countdown(writer io.Writer) {
	for i := countdownStart; i >= 1; i-- {
		fmt.Fprintln(writer, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(writer, finalWord)
}
