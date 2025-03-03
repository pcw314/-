package api

import (
	"fmt"
	"gitee.com/xygfm/authorization/apps/user"
	service "gitee.com/xygfm/authorization/middleware"
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
	var req user.CreatedStaff
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	fmt.Println("42342312", utils.GetUserRole(ctx))
	if utils.GetUserRole(ctx) != 3 {
		response.Error(ctx, result.DefaultError("无权限"))
		return
	}
	if req.Username == "" || req.Password == "" {
		response.Error(ctx, result.DefaultError("缺少参数"))
		return
	} else {
		req.Password = utils.BcryptHash(req.Password)
	}

	u, err := h.svc.Register(ctx, &user.User{
		Username: req.Username,
		Password: req.Password,
		Role:     3,
		State:    1,
	})
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}

	_, err = h.svc.CreateStaff(ctx, &user.Staff{
		UserID: u.ID,
		Name:   req.Name,
		Phone:  req.Phone,
		Avatar: req.Avatar,
	})
	if err != nil {
		err = h.svc.DeleteStaff(ctx, u.ID)
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("创建员工成功", ""))
	return
}

func (h *handler) CreateStudent(ctx *gin.Context) {
	var req user.CreatedStudent
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	var userInfo user.User
	userInfo.Username = req.Username
	userInfo.Password = utils.BcryptHash(req.Password)
	userInfo.Role = 1
	userInfo.State = 1
	u, err := h.svc.Register(ctx, &userInfo)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}

	_, err = h.svc.CreateStudent(ctx, &user.Student{
		UserID:   u.ID,
		Name:     req.Name,
		Sex:      req.Sex,
		Age:      req.Age,
		Major:    req.Major,
		Phone:    req.Phone,
		SchoolID: req.SchoolID,
		Avatar:   req.Avatar,
	})
	if err != nil {
		err = h.svc.DeleteStaff(ctx, u.ID)
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("创建学生成功", ""))
	return
}
func (h *handler) CreateEnterprise(ctx *gin.Context) {
	var req user.CreatedEnterprise
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	userInfo := user.User{
		Username: req.Username,
		Password: utils.BcryptHash(req.Password),
		Role:     2,
		State:    1,
	}
	u, err := h.svc.Register(ctx, &userInfo)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	_, err = h.svc.CreateEnterprise(ctx, &user.Enterprise{
		UserID:     u.ID,
		Name:       req.Name,
		LegaPerson: req.LegaPerson,
		Phone:      req.Phone,
		SchoolID:   req.SchoolID,
		Avatar:     req.Avatar,
	})
	if err != nil {
		err = h.svc.DeleteStaff(ctx, u.ID)
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("创建商户成功", ""))
	return
}
func (h *handler) Register(ctx *gin.Context) {
	var req user.User
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	if req.Username == "" || req.Password == "" || req.Role == 0 {
		response.Error(ctx, result.DefaultError("缺少参数"))
		return
	} else {
		req.Password = utils.BcryptHash(req.Password)
	}
	u, err := h.svc.Register(ctx, &req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	if req.Role == 1 {
		_, err = h.svc.CreateStudent(ctx, &user.Student{UserID: u.ID})
		if err != nil {
			response.Error(ctx, result.DefaultError(err.Error()))
			return
		}
	} else if req.Role == 2 {
		_, err = h.svc.CreateEnterprise(ctx, &user.Enterprise{UserID: u.ID})
		if err != nil {
			response.Error(ctx, result.DefaultError(err.Error()))
			return
		}
	} else {
		response.Error(ctx, result.DefaultError("参数错误"))
		return
	}
	token, err := service.MakeToken(req.ID, req.Username, req.Role)
	response.Success(ctx, result.NewCorrect("注册成功", token))
	return

}

func (h *handler) ListMenu(ctx *gin.Context) {
	fmt.Println("进入获取菜单的api层")
	fmt.Println("ru=========", utils.GetUserID(ctx))
	var ids = []int{1, 2, 3}
	menu, err := h.svc.ListMenu(ctx, &ids)
	if err != nil {
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
		response.Success(ctx, result.NewCorrect("获取成功", response.Paging{
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
	//if utils.GetUserID(ctx) != 1 && utils.GetUserID(ctx) != id {
	//	response.Error(ctx, result.DefaultError("无权限"))
	//}
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

func (h *handler) UpdateStaff(ctx *gin.Context) {
	id := cast.ToInt(ctx.Param("id"))
	var staff *user.CreatedStaff
	err := ctx.ShouldBindJSON(&staff)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	staff.ID = id
	_, err = h.svc.UpdateStaff(ctx, staff)
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

func (h *handler) InitializePassword(ctx *gin.Context) {
	id := cast.ToInt(ctx.Param("id"))
	if utils.GetUserRole(ctx) != 3 {
		response.Error(ctx, result.DefaultError("错误"))
		return
	}
	err := h.svc.InitializePassword(ctx, id)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("密码初始化成功", ""))
}

func (h *handler) UpdateUserState(ctx *gin.Context) {
	id := cast.ToInt(ctx.Param("id"))
	if utils.GetUserRole(ctx) != 3 {
		response.Error(ctx, result.DefaultError("错误"))
		return
	}
	err := h.svc.UpdateUserState(ctx, id)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("状态修改成功", ""))
}
