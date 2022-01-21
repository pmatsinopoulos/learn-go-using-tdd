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

func (dict Dictionary) Add(word, explanation string) (err error) {
	if _, err := dict.Search(word); err == nil {
		return errors.New("word already present in dictionary")
	}

	dict.dict[word] = explanation
	err = nil
	return
}
