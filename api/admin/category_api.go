package admin

import (
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
}

func NewCategoryController() *CategoryController {
	return &CategoryController{}
}

// CategoryPages 分页获取分类数据
func (category *CategoryController) CategoryPages(c *gin.Context) {

}

// CategoryList 获取所有的分类信息
func (category *CategoryController) CategoryList(c *gin.Context) {

}

// CategoryAdd 分类增加
func (category *CategoryController) CategoryAdd(c *gin.Context) {

}

// CategoryUpdate 分类增加
func (category *CategoryController) CategoryUpdate(c *gin.Context) {

}

// CategoryDel 删除（支持批量）
func (category *CategoryController) CategoryDel(c *gin.Context) {

}
