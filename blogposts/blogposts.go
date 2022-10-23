package blogposts

import (
	"io/fs"
)

func NewPostsFromFS(filesystem fs.FS) ([]Post, error) {
	dirEntries, err := fs.ReadDir(filesystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, dirEntry := range dirEntries {
		post, err := getPost(filesystem, dirEntry.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}
