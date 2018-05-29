package main

import "testing"

func TestDecode(t *testing.T) {
	post, err := decode("post.json")
	if err != nil {
		t.Error(err)
	}
	if post.Id != 1 {
		t.Error("Wrong id, was expecting 1 but got", post.Id)
	}
	if post.Content != "Hello World!" {
		t.Error("Wrong content, was expecting ", post.Content)
	}
}
func testEncode(t *testing.T) {
	t.Skip("Skipping encoding for now")
}
func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decode("post.json")
	}
}