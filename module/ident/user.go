package ident

type User struct {
	UserName string `bson:"user_name"`
	Passwd   string `bson:"password"`
	Id       int    `bson:"id"`
}
