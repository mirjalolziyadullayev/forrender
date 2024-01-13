package models

type UserModel struct {
	ID int
	Firstname string
	Lastname string
	Posts []PostModel
}