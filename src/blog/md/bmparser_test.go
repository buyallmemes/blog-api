package md

import (
	"reflect"
	"testing"
)

func TestSimpleMdToHTMLParsing(t *testing.T) {
	actual := ToHTML("Hello, World!")

	expected := &ParserMD{
		Title:  "",
		Html:   "<p>Hello, World!</p>\n",
		Anchor: "",
		Date:   "",
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected:\n%#v\nbut was:\n%#v\n", expected, actual)
	}
}

func TestFrontMatterMdToHTMLParsing(t *testing.T) {
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
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected:\n%#v\nbut was:\n%#v\n", expected, actual)
	}
}
