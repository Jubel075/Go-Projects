package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts/1"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	var post Post
	if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Printf("Post #%d\nTitle: %s\nBody: %s\n", post.ID, post.Title, post.Body)
}
