package oss

import (
	"gitee.com/xygfm/authorization/response"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type Service interface {
	// ListFile 获取文件列表
	ListFile(ctx *gin.Context, req *response.Paging) ([]*File, int64, error)
	// CreateFile 创建文件记录
	CreateFile(ctx *gin.Context, file *File) error
	// UploadFiles 文件上传
	UploadFiles(ctx *gin.Context, files []*multipart.FileHeader) ([]File, error)
}

type ServiceImpl struct{}

func (s *ServiceImpl) ListFile(ctx *gin.Context, req *response.Paging) ([]*File, int64, error) {
	return nil, 0, nil
}
func (s *ServiceImpl) CreateFile(ctx *gin.Context, file *File) error {
	return nil
}
func (s *ServiceImpl) UploadFiles(ctx *gin.Context, files []*multipart.FileHeader) ([]File, error) {
	return []File{}, nil
}
