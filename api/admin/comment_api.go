package admin

import (
	"github.com/gin-gonic/gin"
)

type CommentController struct {
}

func NewCommentController() *CommentController {
	return &CommentController{}
}

// CommentPages 评论信息分页展示
func (comm *CommentController) CommentPages(c *gin.Context) {

}

// CommentDelete 删除评论（支持批量删除）
func (comm *CommentController) CommentDelete(c *gin.Context) {

}
