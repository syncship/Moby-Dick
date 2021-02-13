package models

// UserModel is the model for a user
// and should not be instantiated directly (use NewUserModel instead)
type UserModel struct {
	ID       int `storm:"id"`
	Warnings int `storm:"index"`
}

// NewUserModel returns an instance of UserModel
func NewUserModel(userID int) UserModel {
	return UserModel{
		ID:       userID,
		Warnings: 0,
	}
}
