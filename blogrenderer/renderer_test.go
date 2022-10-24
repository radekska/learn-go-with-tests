package blogrenderer_test

import (
	"bytes"
	approvals "github.com/approvals/go-approval-tests"
	"io"
	"learn-go-with-tests/blogposts"
	"learn-go-with-tests/blogrenderer"
	"testing"
)

func TestRenderer(t *testing.T) {
	var aPost = blogposts.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}
	renderer, err := blogrenderer.NewPostRender()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := renderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = blogposts.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	renderer, err := blogrenderer.NewPostRender()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		renderer.Render(io.Discard, aPost)
	}
}

// TODO - https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/html-templates#back-to-the-real-work
