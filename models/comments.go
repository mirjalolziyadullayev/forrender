package models

type CommentModel struct {
	ID int
	UserID int
	PostID int
	Content string
}