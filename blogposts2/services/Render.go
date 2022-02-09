package services

import (
	"embed"
	"github.com/pmatsinopoulos/blogposts2/models"
	"html/template"
	"io"
)

type PostRenderer struct {
	templ *template.Template
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

// Writes data to the buffer. The data should be the HTML representation
// of the +Post+

func (r *PostRenderer) Render(buf io.Writer, post models.Post) (err error) {
	err = r.templ.Execute(buf, post)

	return err
}
