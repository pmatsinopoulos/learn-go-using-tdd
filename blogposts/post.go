package blogposts

import (
	"bufio"
	"io"
)

type Post struct {
	Description string
	Title       string
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)
	readLine := func() string {
		scanner.Scan()
		return scanner.Text()
	}
	title := readLine()[7:]
	description := readLine()[13:]

	post := Post{
		Description: description,
		Title:       title,
	}
	return post, nil

}
