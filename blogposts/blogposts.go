package blogposts

import (
	"io/fs"
)

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, dirEntry := range dir {
		post, err := getPost(fileSystem, dirEntry.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, dirEntry string) (Post, error) {
	postFile, err := fileSystem.Open(dirEntry)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}
