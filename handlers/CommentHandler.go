package handlers

import "net/http"

func CommentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetAllComments(w,r)
	case "POST":
		CreateComment(w,r)
	case "PUT":
		UpdateComment(w,r)
	case "DELETE":
		DeleteComment(w,r)
	}
}
func CreateComment(w http.ResponseWriter, r *http.Request) {
	
}
func UpdateComment(w http.ResponseWriter, r *http.Request) {
	
}
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	
}
func GetAllComments(w http.ResponseWriter, r *http.Request) {
	
}