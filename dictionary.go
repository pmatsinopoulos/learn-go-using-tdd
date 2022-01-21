package main

import "errors"

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
