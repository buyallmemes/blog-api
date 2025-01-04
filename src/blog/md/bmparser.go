package md

import (
	"bytes"
	"github.com/gosimple/slug"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/frontmatter"
	"log"
)

type ParserMD struct {
	Html   string
	Title  string
	Date   string
	Anchor string
}

func ToHTML(source string) *ParserMD {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			&frontmatter.Extender{},
		),
	)
	var buf bytes.Buffer
	context := parser.NewContext()
	withContext := parser.WithContext(context)
	err := markdown.Convert([]byte(source), &buf, withContext)
	if err != nil {
		log.Fatal(err)
	}

	md := ParserMD{
		Html: buf.String(),
	}
	d := frontmatter.Get(context)
	if d != nil {
		var meta struct {
			Title string `yaml:"title"`
			Date  string `yaml:"date"`
		}

		err = d.Decode(&meta)
		if err != nil {
			log.Fatal(err)
		}
		if title := &meta.Title; title != nil {
			md.Title = *title
			md.Anchor = slug.Make(*title)
		}
		if date := &meta.Date; date != nil {
			md.Date = *date
		}
	}

	return &md
}
