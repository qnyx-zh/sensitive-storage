package entity

type Password struct {
	BaseField   `gorm:"embedded"` // 公有字段
	Topic       string            `gorm:"column:topic;type:string;size:1000" json:"topic"`             // 标题
	UserId      uint              `gorm:"column:user_id;type:integer;size:1000" json:"userId"`         // 所属人id
	Username    string            `gorm:"column:username;type:string;size:1000" json:"username"`       // 账户
	Password    string            `gorm:"column:password;type:string;size:1000" json:"password"`       // 密码
	Description string            `gorm:"column:description;type:string;size:1000" json:"description"` // 备注
}
