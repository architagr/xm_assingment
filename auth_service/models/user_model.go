package models

type UserModel struct {
	Id       interface{}   `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	UserName string        `json:"userName" bson:"userName"`
	Password PasswordModel `json:"-" bson:"password"`
}
