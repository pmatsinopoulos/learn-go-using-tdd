package services

import (
	"bytes"
	"fmt"
	"github.com/pmatsinopoulos/blogposts2/models"
)

// Writes data to the buffer. The data should be the HTML representation
// of the +Post+

func Render(buf *bytes.Buffer, post models.Post) (err error) {
	template := `<h1>%s</h1>
<p>%s</p>
Tags: <ul><li>%s</li><li>%s</li></ul>`
	_, err = fmt.Fprintf(buf, template, post.Title, post.Description, post.Tags[0], post.Tags[1])

	return err
}
