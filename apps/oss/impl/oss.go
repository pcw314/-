package impl

import (
	"fmt"
	"gitee.com/xygfm/authorization/apps/oss"
	"gitee.com/xygfm/authorization/response"
	utils "gitee.com/xygfm/authorization/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func (i *impl) ListFile(ctx *gin.Context, req *response.Paging) ([]*oss.File, int64, error) {
	fmt.Println("业务层")
	limit := req.Size
	offset := (req.Page - 1) * limit
	var fileList []*oss.File
	db := i.mdb.Model(&oss.File{})
	if req.Search != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Search))
	}
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Order("id desc").Find(&fileList).Error
	if err != nil {
		return nil, 0, err
	}

	return fileList, total, nil
}

func (i *impl) CreateFile(ctx *gin.Context, req *oss.File) error {
	err := i.mdb.Model(&oss.File{}).Create(&req).Error
	if err != nil {
		return err
	}
	return nil
}

func (i *impl) UploadFiles(ctx *gin.Context, files []*multipart.FileHeader) ([]oss.File, error) {
	var fileInfos []oss.File

	for _, file := range files {
		// 生成文件保存路径
		now := time.Now()
		savePath := filepath.Join("storage/uploads",
			strconv.Itoa(utils.GetUserRole(ctx)),
			strconv.Itoa(utils.GetUserID(ctx)),
			strconv.Itoa(now.Year()),
			fmt.Sprintf("%02d", now.Month()),
			fmt.Sprintf("%02d", now.Day()))
		// 确保目录存在
		if err := os.MkdirAll(savePath, 0755); err != nil {
			return nil, fmt.Errorf("创建目录失败: %v", err)
		}

		// 生成唯一文件名
		ext := filepath.Ext(file.Filename)
		fileName := uuid.NewString() + ext
		filePath := filepath.Join(savePath, fileName)

		// 保存文件
		if err := ctx.SaveUploadedFile(file, filePath); err != nil {
			return nil, fmt.Errorf("保存文件失败: %v", err)
		}
		fileURL := fmt.Sprintf("/%s", filepath.ToSlash(filePath)[len("storage/"):])
		// 创建文件记录
		fileInfo := oss.File{
			Name:      strings.TrimSuffix(file.Filename, filepath.Ext(file.Filename)), // 原始文件名
			Path:      fileURL,                                                        // 保存路径
			Size:      strconv.FormatInt(file.Size, 10),                               // 文件大小
			Type:      utils.GetFileType(file.Filename),                               // 文件类型
			OwnerType: strconv.Itoa(utils.GetUserRole(ctx)),
			OwnerID:   utils.GetUserID(ctx),
			CreatedAt: time.Now().UnixMilli(),
		}
		err := i.CreateFile(ctx, &fileInfo)
		if err != nil {
			return nil, err
		}
		fileInfo.Path = utils.GetApiByType("file").Url + fileInfo.Path
		fileInfos = append(fileInfos, fileInfo)
	}

	return fileInfos, nil
}
