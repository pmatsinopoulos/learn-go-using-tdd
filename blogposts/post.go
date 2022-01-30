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
	scanner.Scan()
	titleLine := scanner.Text()

	scanner.Scan()
	descriptionLine := scanner.Text()

	post := Post{
		Description: descriptionLine[13:],
		Title:       titleLine[7:],
	}
	return post, nil

}
