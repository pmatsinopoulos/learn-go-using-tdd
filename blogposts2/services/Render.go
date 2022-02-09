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
Tags: <ul>`
	for _, tag := range post.Tags {
		template = fmt.Sprintf("%s<li>%s</li>", template, tag)
	}
	template = fmt.Sprintf("%s</ul>", template)
	_, err = fmt.Fprintf(buf, template, post.Title, post.Description)

	return err
}
