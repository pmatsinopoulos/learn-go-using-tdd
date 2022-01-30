package property_based_tests

import (
	"fmt"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	testCases := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
		11: "XI",
		12: "XII",
		13: "XIII",
		14: "XIV",
		15: "XV",
		16: "XVI",
		17: "XVII",
		18: "XVIII",
		19: "XIX",
		20: "XX",
		39: "XXXIX",
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
