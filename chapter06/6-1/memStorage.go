package main

import "fmt"

type Post struct {
	Id int
	Content string
	Author string
}

var PostById map[int]*Post
var PostByAuthor map[string][]*Post

func store(post Post) {
	PostById[post.Id] = &post
	PostByAuthor[post.Author] = append(PostByAuthor[post.Author], &post)
}

func main() {
	PostById = make(map[int]*Post)
	PostByAuthor = make(map[string][]*Post)

	post1 := Post{Id: 1, Content: "Hello World", Author: "pandoraemon"}
	post2 := Post{Id: 2, Content: "Bonjour Monde!", Author: "Dora"}
	post3 := Post{3, "你好 世界!", "潘多拉"}
	post4 := Post{4, "Greetings Earthlings", "Dora"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])


	for _, post := range PostByAuthor["Dora"] {
		fmt.Println(post)
	}

	for _, post := range PostByAuthor["潘多拉"] {
		fmt.Println(post)
	}
}
