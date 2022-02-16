package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/rinwaowuogba/blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: go, python
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, java
---
B
L
M`
	)

	fs := fstest.MapFS{
		"hello world.md": {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	assertPost(t, posts[0], blogposts.Post{
		Title: "Post 1", Description: "Description 1",
		Tags: []string{
			"go",
			"python",
		},
		Body: `Hello
World`,
	})
}

func assertPost(t *testing.T, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got,want)
	}
}