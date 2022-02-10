package services

import (
	"embed"
	"github.com/pmatsinopoulos/blogposts2/models"
	"github.com/pmatsinopoulos/blogposts2/viewModels"
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
	err = r.templ.ExecuteTemplate(buf, "blog.gohtml", post)

	return err
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []models.Post) error {
	postViews := make([]viewModels.PostViewModel, 0, len(posts))
	for _, post := range posts {
		postViews = append(postViews, viewModels.PostViewModel{Post: post})
	}

	if err := r.templ.ExecuteTemplate(w, "index.gohtml", postViews); err != nil {
		return err
	}

	return nil
}
