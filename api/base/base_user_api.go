// Package base
/**
 * @Author: r0
 * @Mail: boogieLing_o@qq.com
 * @Description: 用户相关操作的api
 * @File:  base_user_api
 * @Version: 1.0.0
 * @Date: 2022/7/4 15:02
 */
package base

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"r0Website-server/models/views"
	"r0Website-server/service"
	"r0Website-server/utils/msg"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{service.NewUserService()}
}

// Login 用户登录
func (u *UserController) Login(c *gin.Context) {
	var params views.LoginVo
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, msg.NewMsg().Failed("查询参数异常"))
		return
	}
	Login, err := u.userService.UserLogin(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, msg.NewMsg().Failed(err.Error()))
		return
	}
	c.JSON(http.StatusOK, msg.NewMsg().Success(Login))
}

// Register 用户注册
func (u *UserController) Register(c *gin.Context) {
	var params views.RegisterVo
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, msg.NewMsg().Failed("查询参数异常"))
		return
	}
	register, err := u.userService.UserRegister(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, msg.NewMsg().Failed(err.Error()))
		return
	}
	c.JSON(http.StatusOK, msg.NewMsg().Success(register))
}
