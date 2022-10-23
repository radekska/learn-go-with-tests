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
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondPost = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
Body of the
second post
here`
	)

	files := fstest.MapFS{
		"hello-world-1.md": {Data: []byte(firstPost)},
		"hello-world-2.md": {Data: []byte(secondPost)},
	}

	posts, _ := blogposts.NewPostsFromFS(files)

	wanted := []blogposts.Post{
		{Title: "Post 1", Description: "Description 1", Tags: []string{"tdd", "go"}, Body: `Hello
World`},
		{Title: "Post 2", Description: "Description 2", Tags: []string{"rust", "borrow-checker"}, Body: `Body of the
second post
here`},
	}

	assert.Equal(t, wanted, posts)
}
