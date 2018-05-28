package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
)

type Post struct {
	Id      int    `json:"id"`
	Content string    `json:"content"`
	Author  Author    `json:"author"`
	Comment []Comment `json:"comments"`
}

type Author struct {
	Id   int `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int `json:"id"`
	Content string `json:"content"`
	Author  Author `json:"author"`
}

func main() {
	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println("Error opening json file", err)
		return
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading json data", err)
		return
	}
	var post Post
	json.Unmarshal(jsonData, &post)
	fmt.Println(post)

	post = Post {
		Id: 1,
		Content: "Hello World!",
		Author: Author {
			2,
			"Sau Sheong",
		},
	}
	output, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling to json:", err)
		return
	}
	err = ioutil.WriteFile("post_write.json", output, 0644)
	if err != nil {
		fmt.Println("Error writing XML to file:", err)
		return
	}

}
