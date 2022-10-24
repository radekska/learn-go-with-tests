package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"learn-go-with-tests/blogposts"
)

// With embed, the files are included in your Go program when you build it.
// This means once you've built your program (which you should only do once),
// the files are always available to you. --> SO COOL!

// Package embed provides access to files embedded in the running Go program.
// Go source files that import "embed" can use the
// go:embed directive to initialize a variable of type string, []byte,
// or FS with the contents of files read from the package directory
// or subdirectories at compile time.

//go:embed "templates/*"
var postTemplates embed.FS

type PostRenderer struct {
	templ *template.Template
}

func NewPostRender() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(writer io.Writer, post blogposts.Post) error {
	if err := r.templ.Execute(writer, post); err != nil {
		return err
	}
	return nil
}
