package main

type Dictionary struct {
	dict map[string]string
}

func (dict Dictionary) Search(word string) string {
	return dict.dict[word]
}
