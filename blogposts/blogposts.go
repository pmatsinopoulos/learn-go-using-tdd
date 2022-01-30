package blogposts

import (
	"io/fs"
)

type Post struct {
}

func NewPostsFromFS(fileSystem fs.FS) (posts []Post, error error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		error = err
		return
	}
	for range dir {
		posts = append(posts, Post{})
	}
	return
}
