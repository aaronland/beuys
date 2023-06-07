package beuys

import (
	"context"
	"fmt"

	"github.com/anaskhan96/soup"
)

const BLOG string = "https://felt.com/blog/"

type Post struct {
	Title string
	Link  string
}

func Posts(ctx context.Context) ([]*Post, error) {

	rsp, err := soup.Get(BLOG)

	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve %s, %w", BLOG, err)
	}

	doc := soup.HTMLParse(rsp)

	links := doc.FindAll("a", "class", "blog-link")
	posts := make([]*Post, len(links))

	for idx, l := range links {

		title := l.FullText()
		attrs := l.Attrs()

		link := fmt.Sprintf("https://felt.com%s", attrs["href"])
		
		post := &Post{
			Title: title,
			Link:  link,
		}

		posts[idx] = post
	}

	return posts, nil
}
