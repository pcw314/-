package api

import (
	"fmt"
	"gitee.com/xygfm/authorization/response"
	"gitee.com/xygfm/authorization/response/result"
	"github.com/gin-gonic/gin"
)

func (h *handler) List(ctx *gin.Context) {
	fmt.Println("我是Api层")
	var req response.Paging
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	if list, total, err := h.svc.ListFile(ctx, &req); err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	} else {
		response.Success(ctx, result.NewCorrect("获取成功", response.Paging{
			Size:   req.Size,
			Page:   req.Page,
			List:   list,
			Total:  total,
			Search: req.Search,
		}))
	}

}

func (h *handler) UploadFiles(c *gin.Context) {
	// 获取上传的文件
	form, err := c.MultipartForm()
	fmt.Println("=============")
	fmt.Println(form)
	if err != nil {
		fmt.Println(err)
		fmt.Println("11111111111")
		response.Error(c, result.DefaultError("获取文件失败"))
		return
	}
	files := form.File["files"]

	// 检查文件数量
	if len(files) == 0 {
		fmt.Println(err)
		fmt.Println("2222222222222222")
		response.Error(c, result.DefaultError("请选择文件"))
		return
	}

	fmt.Println("111-----------------1111")
	// 调用服务层处理文件上传
	fileInfos, err := h.svc.UploadFiles(c, files)
	if err != nil {
		response.Error(c, result.DefaultError(err.Error()))
		return
	}

	response.Success(c, result.NewCorrect("上传成功", fileInfos))
}

func (h *handler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.svc.DeleteFile(ctx, id); err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("删除成功", ""))

}
