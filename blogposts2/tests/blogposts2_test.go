package tests

import (
	"bytes"
	approvals "github.com/approvals/go-approval-tests"
	"github.com/pmatsinopoulos/blogposts2/models"
	"github.com/pmatsinopoulos/blogposts2/services"
	"testing"
)

func TestRenderer(t *testing.T) {
	var (
		aPost = models.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := services.Render(&buf, aPost)
		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}
