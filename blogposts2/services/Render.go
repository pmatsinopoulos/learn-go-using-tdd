package services

import (
	"bytes"
	"fmt"
	"github.com/pmatsinopoulos/blogposts2/models"
)

// Writes data to the buffer. The data should be the HTML representation
// of the +Post+

func Render(buf *bytes.Buffer, post models.Post) (err error) {
	_, err = fmt.Fprintf(buf, "<h1>%s</h1>", post.Title)

	return err
}
