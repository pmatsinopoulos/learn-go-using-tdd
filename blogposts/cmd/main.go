package main

import (
	"githug.com/pmatsinopoulos/blogposts"
	"log"
	"os"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	for _, post := range posts {
		log.Println(post)
	}
}
