package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	title := readMetaLine(titleSeparator)
	desc := readMetaLine(descriptionSeparator)
	tags := strings.Split(readMetaLine(tagsSeparator), ", ")

	scanner.Scan()

	buf := bytes.Buffer{}

	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}

	body := strings.TrimSuffix(buf.String(), "\n")

	return Post{
		Title:       title,
		Description: desc,
		Tags:        tags,
		Body:        body,
	}, nil
}
