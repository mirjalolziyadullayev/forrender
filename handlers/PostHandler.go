package handlers

import (
	"encoding/json"
	"fmt"
	"instagram/helper"
	"instagram/models"
	"net/http"
	"os"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetAllPosts(w, r)
	case "POST":
		CreatePost(w, r)
	case "PUT":
		UpdatePost(w, r)
	case "DELETE":
		DeletePost(w, r)
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var newPost models.PostModel
	json.NewDecoder(r.Body).Decode(&newPost)

	var PostsData []models.PostModel
	PostByte, _ := os.ReadFile("db/posts.json")
	json.Unmarshal(PostByte, &PostsData)

	//-----------------------------------

	var UsersData []models.UserModel
	UserByte, _ := os.ReadFile("db/users.json")
	json.Unmarshal(UserByte, &UsersData)

	var userFound bool
	for i := 0; i < len(UsersData); i++ {
		if UsersData[i].ID == newPost.UserID {
			userFound = true
			break
		}
	}
	if !userFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "User not found with ID", newPost.UserID)
		return
	}
	newPost.ID = helper.MaxIDPost(PostsData)
	PostsData = append(PostsData, newPost)

	res, _ := json.Marshal(PostsData)
	os.WriteFile("db/posts.json", res, 0)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Post Created!")
	fmt.Println("Post Created")
}
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	var updatePost models.PostModel
	json.NewDecoder(r.Body).Decode(&updatePost)

	var PostsData []models.PostModel
	PostByte, _ := os.ReadFile("db/posts.json")
	json.Unmarshal(PostByte, &PostsData)

	//-----------------------------------

	var UsersData []models.UserModel
	UserByte, _ := os.ReadFile("db/users.json")
	json.Unmarshal(UserByte, &UsersData)

	var userFound bool
	for i := 0; i < len(UsersData); i++ {
		if UsersData[i].ID == updatePost.UserID {
			userFound = true
			break
		}
	}
	if !userFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "User not found with ID", updatePost.UserID)
		return
	}

	var postFound bool
	for i := 0; i < len(PostsData); i++ {
		if PostsData[i].ID == updatePost.ID {
			if updatePost.Title != "" {
				PostsData[i].Title = updatePost.Title
			}
			if updatePost.Content != "" {
				PostsData[i].Content = updatePost.Content
			}
			postFound = true
			break
		}
	}
	if !postFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Post not found with ID", updatePost.ID)
		return
	}

	res, _ := json.Marshal(PostsData)
	os.WriteFile("db/posts.json", res, 0)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Post Created!")
	fmt.Println("Post Created")
}
func DeletePost(w http.ResponseWriter, r *http.Request) {
	var deletePost models.PostModel
	json.NewDecoder(r.Body).Decode(&deletePost)

	var PostsData []models.PostModel
	PostByte, _ := os.ReadFile("db/posts.json")
	json.Unmarshal(PostByte, &PostsData)

	//-----------------------------------

	var UsersData []models.UserModel
	UserByte, _ := os.ReadFile("db/users.json")
	json.Unmarshal(UserByte, &UsersData)

	var userFound bool
	for i := 0; i < len(UsersData); i++ {
		if UsersData[i].ID == deletePost.UserID {
			userFound = true
			break
		}
	}
	if !userFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "User not found with ID", deletePost.UserID)
		return
	}

	var postFound bool
	for i := 0; i < len(PostsData); i++ {
		if PostsData[i].ID == deletePost.ID {
			PostsData = append(PostsData[:i], PostsData[i+1:]...)
			postFound = true
			break
		}
	}
	if !postFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Post not found with ID", deletePost.ID)
		return
	}

	res, _ := json.Marshal(PostsData)
	os.WriteFile("db/posts.json", res, 0)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Post Deleted!")
	fmt.Println("Post Deleted")
}
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	var PostsData []models.PostModel
	PostByte, _ := os.ReadFile("db/posts.json")
	json.Unmarshal(PostByte, &PostsData)

	var UserData []models.UserModel
	UserByte, _ := os.ReadFile("db/users.json")
	json.Unmarshal(UserByte, &UserData)

	fmt.Fprintln(w, "---------------------------")
	for i := 0; i < len(PostsData); i++ {
		fmt.Fprintln(w, "Post's ID:", PostsData[i].ID)
		fmt.Fprintln(w, "Post's Title:", PostsData[i].Title)
		fmt.Fprintln(w, "Post's Content:", PostsData[i].Content)
		fmt.Fprintln(w, "Post's Likes count:", PostsData[i].Likes)
		fmt.Fprint(w, "----------------------")
	}
	fmt.Fprintln(w, "---------------------------")

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Post Created!")
	fmt.Println("Post Created")
}
