package models

type UserModel struct {
	Model
	Username string `gorm:"size:16" json:"username"`
	Password string `gorm:"size:64" json:"-"`
	Nickname string `gorm:"size:64" json:"nickname"`
	RoleID   int8   `json:"roleID"` // 1 管理员 2 普通用户
}
