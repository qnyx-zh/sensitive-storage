package entity

type User struct {
	BaseField `gorm:"embedded"`
	Username  string `gorm:"column:username;not null;type:string;size:20" json:"username"`
	Password  string `gorm:"column:password;not null;type:string;size:20" json:"password" `
}
