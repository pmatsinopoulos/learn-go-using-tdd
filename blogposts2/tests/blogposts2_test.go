package tests

import (
	"bytes"
	approvals "github.com/approvals/go-approval-tests"
	"github.com/pmatsinopoulos/blogposts2/models"
	"github.com/pmatsinopoulos/blogposts2/services"
	"io"
	"testing"
)

func TestRenderer(t *testing.T) {
	postRenderer, err := services.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		var (
			aPost = models.Post{
				Title:       "hello world",
				Body:        "This is a post",
				Description: "This is a description",
				Tags:        []string{"go", "tdd"},
			}
		)

		buf := bytes.Buffer{}
		err := postRenderer.Render(&buf, aPost)
		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("it renders an index of products", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []models.Post{
			{Title: "Hello World"},
			{Title: "Hello World 2"},
		}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<ol>
  <li><a href="/post/hello-world">Hello World</a></li>
  <li><a href="/post/hello-world-2">Hello World 2</a></li>
</ol>
`
		if got != want {
			t.Errorf("got \n%q, want\n%q", got, want)
		}
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = models.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := services.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}
