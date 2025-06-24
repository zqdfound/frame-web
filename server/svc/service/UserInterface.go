package userService

import (
	"frame-web/global"
	"frame-web/model/request"
	"frame-web/svc/models"
)

// 分页测试
func GetAllUsersPage(info request.PageInfo) (list []models.User, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	var userList []models.User
	db := global.DB.Model(&models.User{})
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	db = db.Limit(limit).Offset(offset)
	err = db.Find(&userList).Error
	return userList, total, err
}

//
//// 根据条件查询用户列表
//func FindUsersByCondition(condition map[string]interface{}) ([]*models.User, error) {
//	var users []*models.User
//	result := db.DB.Where(condition).Find(&users)
//	if result.Error != nil {
//		zlog.Error("Failed to query users by condition", "error", result.Error)
//		return nil, result.Error
//	}
//	return users, nil
//}
