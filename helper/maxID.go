package helper

import "instagram/models"

func MaxIDUser(UserArray []models.UserModel) int {
	var maxID = 0
	for i := 0; i < len(UserArray); i++ {
		if maxID < UserArray[i].ID {
			maxID=UserArray[i].ID
		}
	}
	return maxID+1
}
func MaxIDPost(PostArray []models.PostModel) int {
	var maxID = 0
	for i := 0; i < len(PostArray); i++ {
		if maxID < PostArray[i].ID {
			maxID=PostArray[i].ID
		}
	}
	return maxID+1
}
func MaxIDComment(CommentArray []models.CommentModel) int {
	var maxID = 0
	for i := 0; i < len(CommentArray); i++ {
		if maxID < CommentArray[i].ID {
			maxID=CommentArray[i].ID
		}
	}
	return maxID+1
}
