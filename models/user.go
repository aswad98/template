package models

type UserData struct {
	UserId *int64  `json:"user-id" db:"user_id"`
	Name   *string `json:"name" db:"name"`
	Email  *string `json:"email" db:"email"`
}

func NewUserData() *UserData {
	return &UserData{
		UserId: new(int64),
		Name:   new(string),
		Email:  new(string),
	}
}
