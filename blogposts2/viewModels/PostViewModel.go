package viewModels

import (
	"github.com/pmatsinopoulos/blogposts2/models"
	"strings"
)

type PostViewModel struct {
	Post models.Post
}

func (pvm PostViewModel) SanitisedTitle() string {
	return strings.ToLower(strings.ReplaceAll(pvm.Post.Title, " ", "-"))
}

func (pvm PostViewModel) Title() string {
	return pvm.Post.Title
}
