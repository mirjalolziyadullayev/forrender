package handlers

import (
	"encoding/json"
	"fmt"
	"instagram/models"
	"net/http"
	"os"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getUser(w, r)
	}
}
func getUser(w http.ResponseWriter, r *http.Request) {
	var toGetUser models.GetUser
	json.NewDecoder(r.Body).Decode(&toGetUser)

	var UserData []models.UserModel
	UserByte, _ := os.ReadFile("db/users.json")
	json.Unmarshal(UserByte, &UserData)

	var PostsData []models.PostModel
	PostByte, _ := os.ReadFile("db/posts.json")
	json.Unmarshal(PostByte, &PostsData)

	fmt.Fprintln(w, "---------------------------")
	for i := 0; i < len(UserData); i++ {
		if UserData[i].ID == toGetUser.ID {
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
}
