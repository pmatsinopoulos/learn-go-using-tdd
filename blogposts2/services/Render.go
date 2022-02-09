package services

import (
	"bytes"
	"embed"
	"github.com/pmatsinopoulos/blogposts2/models"
	"html/template"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

// Writes data to the buffer. The data should be the HTML representation
// of the +Post+

func Render(buf *bytes.Buffer, post models.Post) (err error) {
	var parsed *template.Template
	parsed, err = template.ParseFS(postTemplates, "templates/*.go.html")
	if err != nil {
		return err
	}

	err = parsed.Execute(buf, post)

	return err
}
