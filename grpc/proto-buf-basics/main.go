package main

import (
	"fmt"
	"proto-buf-basics/proto"
)

func main() {
	b := proto.BlogRequest{
		BlogId:  101,
		Foo:     "foo",
		Title:   "blog title",
		Content: "some content",
	}

	fmt.Println(b.GetBlogId())
	fmt.Println(b.GetTitle())
	fmt.Println(b.String())
}
