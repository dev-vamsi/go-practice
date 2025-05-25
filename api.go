package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func FetchAndPrintResponse() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		fmt.Println("Error making api call", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("❌ Read error:", err)
		return
	}

	var posts []Post

	json.Unmarshal(body, &posts)
	fmt.Println(posts)
}
