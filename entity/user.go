package entity

type User struct{
	userName string
	password string
	name string
	age int
	createTime int64
	updateTime int64
}

func (entity *User) NewUser() *User{
	return &User{}
}