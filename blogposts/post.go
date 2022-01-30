package blogposts

import (
	"bufio"
	"io"
	"strings"
)

type Post struct {
	Description string
	Title       string
	Tags        []string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)
	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}
	title := readMetaLine(titleSeparator)
	description := readMetaLine(descriptionSeparator)
	tags := readMetaLine(tagsSeparator)
	tagsArray := strings.Split(tags, ", ")

	post := Post{
		Description: description,
		Title:       title,
		Tags:        tagsArray,
	}
	return post, nil

}
