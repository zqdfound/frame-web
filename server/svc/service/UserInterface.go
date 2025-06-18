package userService

import (
	"frame-web/db"
	"frame-web/svc/models"
	zlog "frame-web/zap"
)

// 查询所有用户列表
func GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	result := db.DB.Find(&users)
	if result.Error != nil {
		zlog.Error("Failed to query users", "error", result.Error)
		// return nil, result.Error
		panic(result.Error)
	}
	return users, nil
}

// 根据条件查询用户列表
func FindUsersByCondition(condition map[string]interface{}) ([]*models.User, error) {
	var users []*models.User
	result := db.DB.Where(condition).Find(&users)
	if result.Error != nil {
		zlog.Error("Failed to query users by condition", "error", result.Error)
		return nil, result.Error
	}
	return users, nil
}
