// Package base
/**
 * @Author: r0
 * @Mail: boogieLing_o@qq.com
 * @Description:
 * @File:  article_route
 * @Version: 1.0.0
 * @Date: 2022/7/5 16:36
 */
package base

import (
	"github.com/gin-gonic/gin"
	"r0Website-server/api/base"
)

func InitBaseArticleRouter(Router *gin.RouterGroup) {
	article := base.NewArticleApi()
	group := Router.Group("article")
	{
		group.GET("", article.ArticleSearch)    // 模糊搜素
		group.GET(":id", article.ArticleSearch) // id精确搜索
	}
}
