package main

import (
    "fmt"
    "testing"
)

var assertStrings = func(t testing.TB, got string, want string) {
    t.Helper()
    if got != want {
        t.Errorf("Expected %q, actual %q", want, got)
    }
}

func TestDictionary(t *testing.T) {
    dictionary := Dictionary{map[string]string{"test": "this is just a test"}}

    t.Run("#Search", func(t *testing.T) {
        t.Run("existing word", func(t *testing.T) {
            got, _ := dictionary.Search("test")
            want := "this is just a test"
            assertStrings(t, got, want)
        })

        t.Run("unknown word", func(t *testing.T) {
            _, err := dictionary.Search("foo")
            if err == nil {
                t.Fatal("should raise an error, but it didn't")
            }
            errString := fmt.Sprintf("%v", err)
            expectedError := "word not found"
            assertStrings(t, errString, expectedError)
        })
    })

    t.Run("#Add", func(t *testing.T) {
        t.Run("new word", func(t *testing.T) {
            dictionary.Add("foo", "this is foo")

            got, _ := dictionary.Search("foo")
            want := "this is foo"

            assertStrings(t, got, want)
        })

        t.Run("existing word", func(t *testing.T) {
            originalValue, err := dictionary.Search("test")
            if err != nil {
                panic("'test' should be present in the dictionary for this test to run")
            }

            err = dictionary.Add("test", "foo test")

            if err == nil {
                t.Errorf("it should raise an error because the word trying to add already exists")
            }

            got, _ := dictionary.Search("test")
            assertStrings(t, got, originalValue)
        })
    })

    t.Run("#Update", func(t *testing.T) {
        t.Run("existing word", func(t *testing.T) {
            _, err := dictionary.Search("test")
            if err != nil {
                panic("'test' should be present in the dictionary for this test to run")
            }

            err = dictionary.Update("test", "new value")

            got, err := dictionary.Search("test")
            want := "new value"

            assertStrings(t, got, want)
        })

        t.Run("non existing word", func(t *testing.T) {
            wordToUpdate := "foo-non-existing"
            _, errSearch := dictionary.Search(wordToUpdate)
            if errSearch == nil {
                panic(fmt.Sprintf("%q should not exist in dictionary for this test to run", wordToUpdate))
            }

            errUpdate := dictionary.Update(wordToUpdate, "new explanation")
            if errUpdate == nil {
                t.Errorf("should raise error when trying to update a word that does not exist")
            }

            value, errSearch2 := dictionary.Search(wordToUpdate)
            if errSearch2 == nil && value == "new explanation" {
                t.Errorf("should not have created the word since it was not originally present")
            }
        })
    })
}
