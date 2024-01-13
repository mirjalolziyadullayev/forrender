package handlers

import (
	"encoding/json"
	"fmt"
	"instagram/helper"
	"instagram/models"
	"net/http"
	"os"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetAllUsers(w, r)
	case "POST":
		CreateUser(w, r)
	case "PUT":
		UpdateUser(w, r)
	case "DELETE":
		DeleteUser(w, r)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.UserModel
	json.NewDecoder(r.Body).Decode(&newUser)

	var UserData []models.UserModel
	UserByte, _ := os.ReadFile("db/users.json")
	json.Unmarshal(UserByte, &UserData)

	newUser.ID = helper.MaxIDUser(UserData)

	UserData = append(UserData, newUser)

	res, _ := json.Marshal(UserData)
	os.WriteFile("db/users.json", res, 0)

	fmt.Println("User yaratildi ID", newUser.ID)
	fmt.Fprint(w, "User Yaratildi:")

	json.NewEncoder(w).Encode(newUser)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser models.UserModel
	json.NewDecoder(r.Body).Decode(&updateUser)
	var UserData []models.UserModel
	UserByte, _ := os.ReadFile("db/users.json")
	json.Unmarshal(UserByte, &UserData)
	var UserFound bool
	for i := 0; i < len(UserData); i++ {
		if UserData[i].ID == updateUser.ID {
			if updateUser.Firstname != "" {
				UserData[i].Firstname = updateUser.Firstname
			}
			if updateUser.Lastname != "" {
				UserData[i].Lastname = updateUser.Lastname
			}
			UserFound = true
			break
		}
	}
	if !UserFound {
		fmt.Println("User Topilmadi ID", updateUser.ID)
		fmt.Fprint(w, "User Topilmadi ID", updateUser.ID)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	res, _ := json.Marshal(UserData)
	os.WriteFile("db/users.json", res, 0)
	fmt.Println("User Yangilandi ID", updateUser.ID)
	fmt.Fprint(w, "User Yangilandi:")
	json.NewEncoder(w).Encode(updateUser)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var DeleteUser models.UserModel
	json.NewDecoder(r.Body).Decode(&DeleteUser)

	var UserData []models.UserModel
	UserByte, _ := os.ReadFile("db/users.json")
	json.Unmarshal(UserByte, &UserData)

	var UserFound bool
	for i := 0; i < len(UserData); i++ {
		if UserData[i].ID == DeleteUser.ID {
			UserData = append(UserData[:i], UserData[i+1:]...)
			UserFound = true
			break
		}
	}
	if !UserFound {
		fmt.Println("User Topilmadi ID", DeleteUser.ID)
		fmt.Fprint(w, "User Topilmadi ID", DeleteUser.ID)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	res, _ := json.Marshal(UserData)
	os.WriteFile("db/users.json", res, 0)

	fmt.Println("User Ochirildi ID", DeleteUser.ID)
	fmt.Fprint(w, "User Ochirildi:")

}
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var UserData []models.UserModel
	UserByte, _ := os.ReadFile("db/users.json")
	json.Unmarshal(UserByte, &UserData)

	var PostsData []models.PostModel
	PostByte, _ := os.ReadFile("db/posts.json")
	json.Unmarshal(PostByte, &PostsData)

	fmt.Fprintln(w, "---------------------------")
	for i := 0; i < len(UserData); i++ {

		fmt.Fprintln(w, "User's ID:", UserData[i].ID)
		fmt.Fprintln(w, "User's Firstname:", UserData[i].Firstname)
		fmt.Fprintln(w, "User's Lastname:", UserData[i].Lastname)
		fmt.Fprintln(w, "User's Posts:")
		fmt.Fprintln(w, "  ----------------------")
		for j := 0; j < len(PostsData); j++ {

			if UserData[i].ID == PostsData[j].UserID {
				fmt.Fprintln(w, "  Post's ID:", PostsData[j].ID)
				fmt.Fprintln(w, "  Post's Title:", PostsData[j].Title)
				fmt.Fprintln(w, "  Post's Content:", PostsData[j].Content)
				fmt.Fprintln(w, "  Post's Likes count:", PostsData[j].Likes)
				fmt.Fprintln(w, "  ----------------------")
			}
		}

		fmt.Fprintln(w, "---------------------------")
	}
}
