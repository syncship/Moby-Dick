package models

// UserModel ..
type UserModel struct {
	ID       int `storm:"id"`
	Warnings int `storm:"index"`
}
