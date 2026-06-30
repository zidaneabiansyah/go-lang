package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

type CreatePostRequest struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

type CreatePostResponse struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

func main() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	fmt.Println("HTTP GET - Basic")
	fmt.Println("---")

	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))

	body, _ := io.ReadAll(resp.Body)
	var post Post
	json.Unmarshal(body, &post)
	fmt.Printf("Post: %s\n\n", post.Title)

	fmt.Println("HTTP GET - Dengan Query Parameters")
	fmt.Println("---")

	baseURL := "https://jsonplaceholder.typicode.com/comments"
	params := url.Values{}
	params.Add("postId", "1")

	reqURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	resp2, err := client.Get(reqURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp2.Body.Close()

	body2, _ := io.ReadAll(resp2.Body)
	var comments []Comment
	json.Unmarshal(body2, &comments)

	fmt.Printf("Ditemukan %d komentar untuk post 1:\n", len(comments))
	for i, c := range comments {
		if i >= 3 {
			fmt.Println("  ... (dipotong)")
			break
		}
		fmt.Printf("  - %s (%s)\n", c.Name, c.Email)
	}
	fmt.Println()

	fmt.Println("HTTP GET - Dengan Custom Headers")
	fmt.Println("---")

	req, _ := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Custom-Header", "belajar-go")
	req.Header.Set("User-Agent", "BelajarGo-Client/1.0")

	resp3, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp3.Body.Close()

	fmt.Printf("Status: %s\n", resp3.Status)
	fmt.Printf("X-Powered-By: %s\n\n", resp3.Header.Get("X-Powered-By"))

	fmt.Println("HTTP POST - Create Data")
	fmt.Println("---")

	newPost := CreatePostRequest{
		Title:  "Belajar Go HTTP Client",
		Body:   "INI ADALAH ISI POST DARI GO CLIENT",
		UserID: 1,
	}

	jsonData, _ := json.Marshal(newPost)

	resp4, err := client.Post(
		"https://jsonplaceholder.typicode.com/posts",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer resp4.Body.Close()

	body4, _ := io.ReadAll(resp4.Body)
	var created CreatePostResponse
	json.Unmarshal(body4, &created)

	fmt.Printf("Status: %s\n", resp4.Status)
	fmt.Printf("Created Post: ID=%d, Title=%s\n\n", created.ID, created.Title)

	fmt.Println("HTTP POST - Dengan http.NewRequest (lebih kontrol)")
	fmt.Println("---")

	payload := map[string]interface{}{
		"title":  "Request Custom",
		"body":   "Body dari custom request",
		"userId": 2,
	}
	payloadBytes, _ := json.Marshal(payload)

	req2, _ := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(payloadBytes))
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("Accept", "application/json")

	resp5, err := client.Do(req2)
	if err != nil {
		log.Fatal(err)
	}
	defer resp5.Body.Close()

	body5, _ := io.ReadAll(resp5.Body)
	var result CreatePostResponse
	json.Unmarshal(body5, &result)

	fmt.Printf("Status: %s\n", resp5.Status)
	fmt.Printf("Result: %+v\n\n", result)

	fmt.Println("HTTP PUT - Update Data")
	fmt.Println("---")

	updateData := map[string]interface{}{
		"id":     1,
		"title":  "Post Updated dari Go",
		"body":   "Ini sudah diupdate",
		"userId": 1,
	}
	updateBytes, _ := json.Marshal(updateData)

	req3, _ := http.NewRequest("PUT", "https://jsonplaceholder.typicode.com/posts/1", bytes.NewBuffer(updateBytes))
	req3.Header.Set("Content-Type", "application/json")

	resp6, err := client.Do(req3)
	if err != nil {
		log.Fatal(err)
	}
	defer resp6.Body.Close()

	body6, _ := io.ReadAll(resp6.Body)
	var updated Post
	json.Unmarshal(body6, &updated)

	fmt.Printf("Status: %s\n", resp6.Status)
	_ = updateBytes
	fmt.Printf("Updated: %s\n\n", updated.Title)

	fmt.Println("HTTP PATCH - Partial Update")
	fmt.Println("---")

	patchData := map[string]interface{}{
		"title": "Title Saja Yang Diupdate",
	}

	patchBytes, _ := json.Marshal(patchData)
	req4, _ := http.NewRequest("PATCH", "https://jsonplaceholder.typicode.com/posts/1", bytes.NewBuffer(patchBytes))
	req4.Header.Set("Content-Type", "application/json")

	resp7, err := client.Do(req4)
	if err != nil {
		log.Fatal(err)
	}
	defer resp7.Body.Close()

	body7, _ := io.ReadAll(resp7.Body)
	var patched Post
	json.Unmarshal(body7, &patched)

	fmt.Printf("Status: %s\n", resp7.Status)
	fmt.Printf("Patched Title: %s\n\n", patched.Title)

	fmt.Println("HTTP DELETE")
	fmt.Println("---")

	req5, _ := http.NewRequest("DELETE", "https://jsonplaceholder.typicode.com/posts/1", nil)

	resp8, err := client.Do(req5)
	if err != nil {
		log.Fatal(err)
	}
	defer resp8.Body.Close()

	fmt.Printf("Status: %s\n", resp8.Status)

	fmt.Println()
	fmt.Println("TIMEOUT HANDLING")
	fmt.Println("---")

	slowClient := &http.Client{
		Timeout: 1 * time.Millisecond, // terlalu singkat
	}

	_, err = slowClient.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Printf("Request gagal (expected): %v\n\n", err)
	}

	fmt.Println("RESPONSE BODY CHECK")
	fmt.Println("---")

	resp9, err := client.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatal(err)
	}
	defer resp9.Body.Close()

	fmt.Printf("Content-Length: %d bytes\n", resp9.ContentLength)

	if resp9.StatusCode == http.StatusOK {
		var p Post
		json.NewDecoder(resp9.Body).Decode(&p)
		fmt.Printf("Post ID %d: %s\n", p.ID, p.Title)
	}

	fmt.Println()
	fmt.Println("TIP: Selalu tutup Response Body (defer resp.Body.Close())")
	fmt.Println("TIP: Gunakan http.Client dengan Timeout untuk production")
}
