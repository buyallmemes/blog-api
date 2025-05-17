package md

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldConvertSimpleMdToHTMLParsing(t *testing.T) {
	actual := ToHTML("Hello, World!")

	expected := &ParserMD{
		Title:  "",
		Html:   "<p>Hello, World!</p>\n",
		Anchor: "",
		Date:   "",
	}
	assert.Equal(t, expected, actual)
}

func TestShouldConvertFrontMatterMdToHTMLParsing(t *testing.T) {
	source := `
---
title: Hello, World!
date: 29.03.2024
---
Hello, World!
`
	actual := ToHTML(source)

	expected := &ParserMD{
		Title:  "Hello, World!",
		Html:   "<p>Hello, World!</p>\n",
		Anchor: "hello-world",
		Date:   "29.03.2024",
	}
	assert.Equal(t, expected, actual)
}
