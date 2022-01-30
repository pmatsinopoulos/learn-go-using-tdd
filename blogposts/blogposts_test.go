package blogposts_test

import (
	"githug.com/pmatsinopoulos/blogposts"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello_world.md":  {Data: []byte("Title: Post 1")},
		"hello_world2.md": {Data: []byte("Title: Post 2")},
	}
	posts, err := blogposts.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("it creates the correct number of posts", func(t *testing.T) {
		if len(posts) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}
	})
	t.Run("each post has the correct title", func(t *testing.T) {
		if posts[0].Title != "Post 1" {
			t.Errorf("Post does not have the expected title. Expected %q, got %q", "Post 1", posts[0].Title)
		}
	})
}
