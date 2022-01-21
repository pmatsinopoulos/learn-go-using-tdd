package main

import (
	"errors"
	"fmt"
)

type Dictionary struct {
	dict map[string]string
}

func (dict Dictionary) Search(word string) (result string, err error) {
	result, ok := dict.dict[word]
	if !ok {
		err = errors.New("word not found")
	}
	return
}

func (dict Dictionary) Add(word, explanation string) (err error) {
	if _, err := dict.Search(word); err == nil {
		return errors.New("word already present in dictionary")
	}

	dict.dict[word] = explanation
	err = nil
	return
}

func (dict Dictionary) Update(word, explanation string) (err error) {
	_, searchError := dict.Search(word)
	if searchError != nil {
		err = errors.New(fmt.Sprintf("word: %q does not exist in the dictionary", word))
		return
	}
	dict.dict[word] = explanation
	err = nil
	return
}
