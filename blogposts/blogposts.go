package blogposts

import (
	"io/fs"
)

type Post struct {
}

func NewPostsFromFS(fileSystem fs.FS) (posts []Post, error error) {
	dir, _ := fs.ReadDir(fileSystem, ".")
	for range dir {
		posts = append(posts, Post{})
	}
	return
}
