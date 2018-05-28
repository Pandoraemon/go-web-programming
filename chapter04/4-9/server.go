package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web Programming</title></head>
<body><h1>Hello World</h1></body>
</html>`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "no souch service!")
}

func headExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://www.bing.com")
	w.WriteHeader(302)
}
func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contect-Type", "application/json")
	post := &Post{
		User:    "padoraemon",
		Threads: []string{"first", "second", "third"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headExample)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}
