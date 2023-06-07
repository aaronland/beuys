package beuys

import (
	"context"
	_ "embed"
	"fmt"
	"text/template"
	"io"
	"time"
	"html"
)

//go:embed feed.xml
var feedTemplate string

type feedVars struct {
	Posts []*Post
	Now   time.Time
}

func Feed(ctx context.Context, wr io.Writer) error {

	posts, err := Posts(ctx)

	if err != nil {
		return fmt.Errorf("Failed to derive posts, %w", err)
	}

	now := time.Now()

	funcs := template.FuncMap{
		"Escape": func(raw string) string {
			return html.EscapeString(raw)
		},
	}
	
	t := template.New("posts").Funcs(funcs)
	
	t, err = t.Parse(feedTemplate)

	if err != nil {
		return fmt.Errorf("Failed to parse feed template, %w", err)
	}

	vars := feedVars{
		Posts: posts,
		Now:   now,
	}

	err = t.Execute(wr, vars)

	if err != nil {
		return fmt.Errorf("Failed to render feed template, %w", err)
	}

	return nil
}
