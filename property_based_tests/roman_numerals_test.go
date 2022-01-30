package property_based_tests

import (
	"fmt"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	testCases := map[int]string{
		1: "I",
		2: "II",
		3: "III",
	}
	for arabicNumber, romanNumber := range testCases {
		testName := fmt.Sprintf("%d gets converted to %q", arabicNumber, romanNumber)
		t.Run(testName, func(t *testing.T) {
			got := ConvertToRoman(arabicNumber)
			want := romanNumber

			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}
