package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Post struct {
	Description string
	Title       string
	Tags        []string
	Body        string
}

func (p Post) String() string {
	return fmt.Sprintf("Title: %s\nDescription: %s\nTags: %v\n---\n%s\n", p.Title, p.Description, p.Tags, p.Body)
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

	scanner.Scan() // ignore this line

	body := readBody(scanner)

	post := Post{
		Description: description,
		Title:       title,
		Tags:        tagsArray,
		Body:        body,
	}
	return post, nil
}

func readBody(scanner *bufio.Scanner) string {
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}
