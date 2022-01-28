package reflection

import (
	"reflect"
	"testing"
)

type TestCase struct {
	Name          string
	Input         interface{}
	ExpectedCalls []string
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []TestCase{
		{
			Name:          "Struct with one string field",
			Input:         struct{ Name string }{Name: "Chris"},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "Struct with two string fields",
			Input: struct {
				Name string
				City string
			}{Name: "Chris", City: "London"},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "Struct with string and integer",
			Input: struct {
				Name string
				Age  int
			}{Name: "Chris", Age: 52},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name:          "Struct with nested fields",
			Input:         Person{Name: "Chris", Profile: Profile{Age: 52, City: "London"}},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name:          "Pointers to things",
			Input:         &Person{Name: "Chris", Profile: Profile{Age: 52, City: "London"}},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "Slices",
			Input: []Profile{
				{Age: 33, City: "London"},
				{Age: 34, City: "Paris"},
			},
			ExpectedCalls: []string{"London", "Paris"},
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.Name, func(t *testing.T) {
			var got []string
			walk(testCase.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, testCase.ExpectedCalls) {
				t.Errorf("got %v, expected %v", got, testCase.ExpectedCalls)
			}
		})
	}
}
