package main

import (
	"github.com/jmoiron/sqlx"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

type Post struct {
	Id       int
	Content  string
	AuthorName  string `db: author`
}

var Db *sqlx.DB

func init() {
	var err error
	Db, err = sqlx.Open("postgres", "host=10.130.10.156 user=gwp dbname=gwp password=123456 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.QueryRowx("select id, content, author from posts limit $1", limit)
	if err != nil {
		panic(err)
	}
	rows.Close()
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	post.Comments = []Comment{}
	err = Db.QueryRow("select id, content, author from posts where id = $1",
		id).Scan(&post.Id, &post.Content, &post.Author)
	rows, err := Db.Query("select id, content, author from comments where post_id = $1", id)
	if err != nil {
		return
	}
	for rows.Next() {
		comment := Comment{Post: &post}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			return
		}
		post.Comments = append(post.Comments, comment)
	}
	rows.Close()
	return
}

func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1",
		post.Id, post.Content, post.Author)
	return
}

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}

func main() {

	post := Post{Content: "Hello World", Author: "pandoraemon"}

	post.Create()

	comment := Comment{Content: "Good post!", Author: "Joe", Post: &post}
	comment.Create()
	post, _ = GetPost(1)
	comment = Comment{Content: "Good post2!", Author: "Joe", Post: &post}
	comment.Create()
	readPost, _ := GetPost(post.Id)

	fmt.Println(readPost)
	fmt.Println(readPost.Comments)
	fmt.Println(readPost.Comments[0].Post)
}
