package models

import (
	"frame-web/model"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type User struct {
	model.COMMON_MODEL
	//ID        uint      `gorm:"primaryKey;column:id;comment:主键ID" json:"id,omitempty"`
	Username string `gorm:"size:255;column:username;not null;uniqueIndex;comment:用户名" json:"username"  binding:"required"`
	Password string `gorm:"size:255;column:password;not null;comment:密码" json:"-"`
	Email    string `gorm:"size:255;column:email;comment:邮箱" json:"email,omitempty"`
	Status   int    `gorm:"type:tinyint;column:status;default:1;comment:状态(1:正常 0:禁用)" json:"status,omitempty"`
	//CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;comment:创建时间" json:"created_at,omitempty"`
	//UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime;comment:更新时间" json:"updated_at,omitempty"`
	//DeletedAt time.Time `gorm:"column:deleted_at;comment:删除时间" json:"updated_at,omitempty"`
}

func (User) TableName() string {
	return "users"
}
