package main

import (
	"fmt"
	"testing"
)

func TestRepeatCharacter(t *testing.T) {
	got := RepeatCharacter('a', 10)
	want := "aaaaaaaaaa"
	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func BenchmarkRepeatCharacter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatCharacter('a', 10)
	}
}

func ExampleRepeatCharacter() {
	stringWithRepeatedCharacter := RepeatCharacter('a', 5)
	fmt.Println(stringWithRepeatedCharacter)
	// Output: aaaaa
}
