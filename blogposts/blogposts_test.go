package blogposts_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"learn-go-with-tests/blogposts"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("i'm here to fail")
}

func TestNewBlogPosts(t *testing.T) {
	const (
		firstPost = `Title: Post 1
Description: Description 1`
		secondPost = `Title: Post 2
Description: Description 2`
	)

	files := fstest.MapFS{
		"hello-world-1.md": {Data: []byte(firstPost)},
		"hello-world-2.md": {Data: []byte(secondPost)},
	}

	posts, _ := blogposts.NewPostsFromFS(files)

	wanted := []blogposts.Post{
		blogposts.Post{Title: "Post 1", Description: "Description 1"},
		blogposts.Post{Title: "Post 2", Description: "Description 2"},
	}

	assert.Equal(t, wanted, posts)
}
