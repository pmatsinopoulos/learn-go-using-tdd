package blogposts

import (
	"bufio"
	"io"
)

type Post struct {
	Description string
	Title       string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)
	readLine := func() string {
		scanner.Scan()
		return scanner.Text()
	}
	title := readLine()[len(titleSeparator):]
	description := readLine()[len(descriptionSeparator):]

	post := Post{
		Description: description,
		Title:       title,
	}
	return post, nil

}
