package blogposts_test

import (
	"learn-go-with-tests/blogposts"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello-world.md":   {Data: []byte("uszanowanko")},
		"hello-world-2.md": {Data: []byte("pepo-hola")},
	}

	posts := blogposts.NewPostsFromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}

// TODO - https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/reading-files#error-handling
