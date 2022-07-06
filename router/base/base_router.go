// Package base
/**
 * @Author: r0
 * @Mail: boogieLing_o@qq.com
 * @Description: 基础接口
 * @File:  base
 * @Version: 1.0.0
 * @Date: 2022/7/3 18:44
 */
package base

import (
	"github.com/gin-gonic/gin"
	baseController "r0Website-server/api/base"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	userController := baseController.NewUserController()
	{
		BaseRouter.POST("login", userController.Login)
		BaseRouter.POST("register", userController.Register)
		InitBaseArticleRouter(BaseRouter)
	}
	return BaseRouter
}
