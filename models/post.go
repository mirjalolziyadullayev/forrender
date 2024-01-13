package models

type PostModel struct {
	ID int
	UserID int
	Title string
	Content string
	Likes int
	Comments []CommentModel 
}