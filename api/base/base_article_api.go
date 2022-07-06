// Package base
/**
 * @Author: r0
 * @Mail: boogieLing_o@qq.com
 * @Description: base的文章api
 * @File:  base_article_api
 * @Version: 1.0.0
 * @Date: 2022/7/5 16:33
 */
package base

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"r0Website-server/models/views"
	"r0Website-server/service"
	"r0Website-server/utils/msg"
)

type ArticleApi struct {
	*service.ArticleService
}

func NewArticleApi() *ArticleApi {
	return &ArticleApi{service.NewArticleService()}
}

// ArticleSearch 模糊搜索文章内容，依赖分词冗杂，允许带空格
// ShouldBindJSON > ShouldBind
func (article *ArticleApi) ArticleSearch(c *gin.Context) {
	var params views.BaseArticleSearchVo
	articleID := c.Param("id")
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, msg.NewMsg().Failed("查询参数异常"))
		return
	}
	result, err := article.ArticleService.ArticleBaseSearch(params, articleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, msg.NewMsg().Failed(err.Error()))
		return
	}
	c.JSON(http.StatusOK, msg.NewMsg().Success(result))
}
