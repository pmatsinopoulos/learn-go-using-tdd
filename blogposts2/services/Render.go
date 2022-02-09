package services

import (
	"bytes"
	"github.com/pmatsinopoulos/blogposts2/models"
	"html/template"
)

// Writes data to the buffer. The data should be the HTML representation
// of the +Post+

func Render(buf *bytes.Buffer, post models.Post) (err error) {
	templateStr := `<h1>{{.Title}}</h1>
<p>{{.Description}}</p>
Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>`

	var parsed *template.Template
	parsed, err = template.New("blog").Parse(templateStr)
	if err != nil {
		return err
	}
	err = parsed.Execute(buf, post)

	return err
}
