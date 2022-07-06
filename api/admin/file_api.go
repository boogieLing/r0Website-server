package admin

import (
	"github.com/gin-gonic/gin"
)

type FileController struct {
}

func NewFileController() *FileController {
	return &FileController{}
}

// UploadFile 文件上传
func (file *FileController) UploadFile(c *gin.Context) {

}

// DownloadFile 文件下载
func (file *FileController) DownloadFile(c *gin.Context) {

}

// FileList 文件列表
func (file *FileController) FileList(c *gin.Context) {

}
