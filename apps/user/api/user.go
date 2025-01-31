package api

import (
	"fmt"
	"gitee.com/xygfm/authorization/apps/user"
	"gitee.com/xygfm/authorization/response"
	"gitee.com/xygfm/authorization/response/result"
	utils "gitee.com/xygfm/authorization/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func (h *handler) Login(ctx *gin.Context) {
	fmt.Println("我是登录Api层")
	token, err := h.svc.Login(ctx, &user.LoginRequest{Username: "", Password: ""})
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
	}
	response.Success(ctx, result.NewCorrect("登录成功", token))
}

func (h *handler) CreateStaff(ctx *gin.Context) {
	var req user.User
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	if utils.GetUserRole(ctx) != 3 {
		response.Error(ctx, result.DefaultError("无权限"))
		return
	}
	_, err = h.svc.Register(ctx, &req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("创建成功", ""))
	return
}

func (h *handler) CreateStudent(ctx *gin.Context) {
	var req user.User
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	u, err := h.svc.Register(ctx, &req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	_, err = h.svc.CreateStudent(ctx, &user.Student{UserID: u.ID})
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("注册成功", ""))
	return

}
func (h *handler) CreateEnterprise(ctx *gin.Context) {
	var req user.User
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	u, err := h.svc.Register(ctx, &req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	_, err = h.svc.CreateEnterprise(ctx, &user.Enterprise{UserID: u.ID})
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
}
func (h *handler) ListMenu(ctx *gin.Context) {
	fmt.Println("进入获取菜单的api层")
	fmt.Println("ru=========", utils.GetUserID(ctx))
	var ids = []int{1, 2, 3}
	menu, err := h.svc.ListMenu(ctx, &ids)
	if err != nil {
		fmt.Println("adjnjksjndksndkjas")
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("获取菜单成功", menu))
	return
}

func (h *handler) ListStudent(ctx *gin.Context) {
	var req response.Paging
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	if list, total, err := h.svc.ListStudent(ctx, &req); err != nil {
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

func (h *handler) ListEnterprise(ctx *gin.Context) {
	fmt.Println("我是Api层")
	var req response.Paging
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	if list, total, err := h.svc.ListEnterprise(ctx, &req); err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	} else {
		response.Success(ctx, result.NewCorrect("登录成功", response.Paging{
			Size:   req.Size,
			Page:   req.Page,
			List:   list,
			Total:  total,
			Search: req.Search,
		}))
	}

}
func (h *handler) ListStaff(ctx *gin.Context) {
	fmt.Println("我是Api层")
	var req response.Paging
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	if list, total, err := h.svc.ListStaff(ctx, &req); err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	} else {
		response.Success(ctx, result.NewCorrect("登录成功", response.Paging{
			Size:   req.Size,
			Page:   req.Page,
			List:   list,
			Total:  total,
			Search: req.Search,
		}))
	}
}

func (h *handler) GetStaff(ctx *gin.Context) {
	id := cast.ToInt(ctx.Param("id"))
	if utils.GetUserID(ctx) != 1 && utils.GetUserID(ctx) != id {
		response.Error(ctx, result.DefaultError("无权限"))
	}
	re, err := h.svc.GetStaffByID(ctx, id)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("获取成功", re))
	return
}
func (h *handler) GetStudent(ctx *gin.Context) {
	id := cast.ToInt(ctx.Param("id"))
	re, err := h.svc.GetStudentByID(ctx, id)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	if utils.GetUserRole(ctx) != 3 && utils.GetUserID(ctx) != re.UserID {
		response.Error(ctx, result.DefaultError("无权限"))
	}
	response.Success(ctx, result.NewCorrect("获取成功", re))
	return
}
func (h *handler) GetEnterprise(ctx *gin.Context) {
	id := cast.ToInt(ctx.Param("id"))
	re, err := h.svc.GetEnterpriseByID(ctx, id)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	if utils.GetUserRole(ctx) != 3 && utils.GetUserID(ctx) != re.UserID {
		response.Error(ctx, result.DefaultError("无权限"))
	}
	response.Success(ctx, result.NewCorrect("获取成功", re))
	return
}

func (h *handler) UpdateStudent(ctx *gin.Context) {
	id := cast.ToInt(ctx.Param("id"))
	var student *user.Student
	err := ctx.ShouldBindJSON(&student)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	student.ID = id
	_, err = h.svc.UpdateStudent(ctx, student)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("修改成功", ""))
}

func (h *handler) UpdateEnterprise(ctx *gin.Context) {
	id := cast.ToInt(ctx.Param("id"))
	var enterprise *user.Enterprise
	err := ctx.ShouldBindJSON(&enterprise)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	enterprise.ID = id
	_, err = h.svc.UpdateEnterprise(ctx, enterprise)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("修改成功", ""))
}

func (h *handler) UpdatePassword(ctx *gin.Context) {
	id := cast.ToInt(ctx.Param("id"))
	var password user.UpdatePasswordRequest
	err := ctx.ShouldBindJSON(&password)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	if utils.GetUserRole(ctx) != 3 && utils.GetUserID(ctx) != id {
		response.Error(ctx, result.DefaultError("错误"))
	}
	err = h.svc.UpdatePassword(ctx, id, password.NewPassword, password.OldPassword)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("密码修改成功", ""))
}

func (h *handler) DeleteStudent(ctx *gin.Context) {
	id := cast.ToInt(ctx.Param("id"))
	err := h.svc.DeleteStudent(ctx, id)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("删除成功", ""))
	return
}

func (h *handler) DeleteEnterprise(ctx *gin.Context) {
	id := cast.ToInt(ctx.Param("id"))
	err := h.svc.DeleteEnterprise(ctx, id)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("删除成功", ""))
}

func (h *handler) DeleteStaff(ctx *gin.Context) {
	id := cast.ToInt(ctx.Param("id"))
	err := h.svc.DeleteStaff(ctx, id)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("删除成功", ""))
}
