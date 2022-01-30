package blogposts_test

import (
	"githug.com/pmatsinopoulos/blogposts"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1`
		secondBody = `Title: Post 2
Description: Description 2`
	)
	fs := fstest.MapFS{
		"hello_world.md":  {Data: []byte(firstBody)},
		"hello_world2.md": {Data: []byte(secondBody)},
	}
	posts, err := blogposts.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}
	assertPost := func(t *testing.T, got blogposts.Post, want blogposts.Post) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v, want %+v", got, want)
		}
	}
	t.Run("it creates the correct number of posts", func(t *testing.T) {
		if len(posts) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}
	})
	assertPost(t, posts[0], blogposts.Post{Title: "Post 1", Description: "Description 1"})
	assertPost(t, posts[1], blogposts.Post{Title: "Post 2", Description: "Description 2"})
}
