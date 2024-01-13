package main

import (
	"fmt"
	"instagram/handlers"
	"net/http"
)

func main() {
	fmt.Println("Server ishlayapti... :8080")

	http.HandleFunc("/user", handlers.UserHandler)
	http.HandleFunc("/post", handlers.PostHandler)
	http.HandleFunc("/comment", handlers.CommentHandler)
	http.HandleFunc("/getuser", handlers.GetUserHandler)

	http.ListenAndServe(":8080", nil)
}