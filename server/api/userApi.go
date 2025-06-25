package api

import (
	"frame-web/global"
	"frame-web/model/request"
	"frame-web/model/response"
	"frame-web/svc/models"
	userService "frame-web/svc/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserApi struct {
}

func (userAPi *UserApi) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	user, err := userService.GetUserById(id)
	if err != nil {
		global.LOG.Error("获取用户失败!", zap.Error(err))
		response.FailWithMessage("获取用户失败:"+err.Error(), c)
		return
	}
	response.OkWithData(user, c)
}

func (userAPi *UserApi) GetUsersPage(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	username := c.Query("username")
	var pageInfo = request.PageInfo{
		Page:     pageNum,
		PageSize: pageSize,
	}
	list, total, err := userService.GetAllUsersPage(&userService.UserPageReq{
		Pages: &pageInfo,
		User: &models.User{
			Username: username,
		},
	})
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (userAPi *UserApi) CreatUsers(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userService.CreateUser(&user); err != nil {
		global.LOG.Error("新增用户失败!", zap.Error(err))
		response.FailWithMessage("新增用户失败:"+err.Error(), c)
		return
	}
	global.LOG.Info("新增用户成功!", zap.Any("user", user))
	response.Ok(c)
}

func (userAPi *UserApi) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	userService.DeleteUserById(id)
	global.LOG.Info("删除用户成功!", zap.Int("id", id))
	response.OkWithMessage("删除成功", c)
}

func (userAPi *UserApi) UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := userService.UpdateUser(&user); err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	global.LOG.Info("更新成功!", zap.Any("user", user))
	response.OkWithMessage("更新成功", c)
}
