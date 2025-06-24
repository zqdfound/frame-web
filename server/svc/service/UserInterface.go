package userService

import (
	"frame-web/global"
	"frame-web/model/request"
	"frame-web/svc/models"
)

// 分页测试
func GetAllUsersPage(info request.PageInfo) (list []models.User, total int64, err error) {
	//limit := info.PageSize
	//offset := info.PageSize * (info.Page - 1)
	var userList []models.User
	db := global.DB.Model(&models.User{})
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	//db = db.Limit(limit).Offset(offset)
	err = db.Scopes(info.Paginate()).Find(&userList).Error
	return userList, total, err
}

// 删除测试
func DeleteUserById(id int) {
	db := global.DB.Model(&models.User{})
	db.Delete(&models.User{}, id)
}

// 更新
func UpdateUser(user *models.User) error {
	return global.DB.Model(&models.User{}).Where("id = ?", user.ID).Updates(&user).Error
}

// 根据id获取用户信息
func GetUserById(id int) (user *models.User, err error) {
	err = global.DB.Where("id = ?", id).First(&user).Error
	return
}

// 新增
func CreateUser(user *models.User) error {
	return global.DB.Create(&user).Error
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
