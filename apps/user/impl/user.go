package impl

import (
	"context"
	"errors"
	"fmt"
	"gitee.com/xygfm/authorization/apps/place"
	"gitee.com/xygfm/authorization/apps/user"
	service "gitee.com/xygfm/authorization/middleware"
	"gitee.com/xygfm/authorization/response"
	utils "gitee.com/xygfm/authorization/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func (i *impl) Login(ctx context.Context, req *user.LoginRequest) (string, error) {
	fmt.Println("我是登录业务层")
	//var name admin.LoginRequest
	//err := global.Mdb.Model(&admin.LoginRequest{}).Find(&name).Error
	//err = i.mdb.Model(&admin.LoginRequest{}).Find(&name).Error
	token, err := service.MakeToken(2, "13979875415", 0)
	if err != nil {
		return "", err
	}
	fmt.Println("name")
	return token, nil
}

// Register 创建用户
func (i *impl) Register(ctx *gin.Context, req *user.User) (*user.User, error) {
	fmt.Println("我是Register业务层")
	err := i.mdb.Model(&user.User{}).
		Where("username = ?", req.Username).
		First(&user.User{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户名已存在")
	}
	err = i.mdb.Model(&user.User{}).Create(&req).Error
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	fmt.Println("req", req)
	return req, nil
}

func (i *impl) CreateStudent(ctx *gin.Context, req *user.Student) (*user.Student, error) {
	err := i.mdb.Model(&user.Student{}).Create(&req).Error
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	return req, nil
}

func (i *impl) CreateEnterprise(ctx *gin.Context, req *user.Enterprise) (*user.Enterprise, error) {
	req.CreatedAt = time.Now().UnixMilli()
	err := i.mdb.Model(&user.Enterprise{}).Create(&req).Error
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	return req, nil
}

func (i *impl) ListMenu(ctx context.Context, req *[]int) ([]*user.MenuRequest, error) {
	fmt.Println("进入获取菜单的业务层")
	// 查询所有菜单数据
	var menus []*user.MenuRequest
	err := i.mdb.Model(&user.MenuRequest{}).Find(&menus).Error
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	// 使用 Map 缓存所有菜单
	menuMap := make(map[int]*user.MenuRequest)
	for _, menu := range menus {
		menuMap[menu.ID] = menu
	}

	// 构建树形结构
	var tree []*user.MenuRequest
	for _, menu := range menus {
		if menu.PID == 0 {
			// 根节点
			tree = append(tree, menu)
		} else {
			// 找到父节点，将当前菜单添加到父节点的 Children 中
			if parent, ok := menuMap[menu.PID]; ok {
				if parent.Children == nil {
					parent.Children = []*user.MenuRequest{}
				}
				parent.Children = append(parent.Children, menu)
			}
		}
	}

	return tree, nil
}

func (i *impl) ListStaff(ctx *gin.Context, req *response.Paging) ([]*user.User, int64, error) {
	if utils.GetUserRole(ctx) != 3 {
		return nil, 0, nil
	}
	limit := req.Size
	offset := (req.Page - 1) * limit
	var pos []*user.User
	db := i.mdb.Model(&user.User{}).Where("role = 3")
	if req.Search != "" {
		db = db.Where("username LIKE ?", fmt.Sprintf("%%%s%%", req.Search))
	}
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Order("id desc").Find(&pos).Error
	if err != nil {
		return nil, 0, err
	}

	return pos, total, nil
}
func (i *impl) ListStudent(ctx *gin.Context, req *response.Paging) ([]*user.Student, int64, error) {
	limit := req.Size
	offset := (req.Page - 1) * limit
	var ids []*int
	db := i.mdb.Model(&user.User{}).Select("id").Where("role = 1")
	if req.Search != "" {
		db = db.Where("username LIKE ?", fmt.Sprintf("%%%s%%", req.Search))
	}
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Order("id desc").Find(&ids).Error
	if err != nil {
		return nil, 0, err
	}
	var pos []*user.Student
	err = i.mdb.Model(&user.Student{}).Where("user_id in (?)", ids).Find(&pos).Error
	if err != nil {
		return nil, 0, err
	}
	return pos, total, nil
}
func (i *impl) ListEnterprise(ctx *gin.Context, req *response.Paging) ([]*user.Enterprise, int64, error) {
	limit := req.Size
	offset := (req.Page - 1) * limit
	var ids []*int
	db := i.mdb.Model(&user.User{}).Select("id").Where("role = 2")
	if req.Search != "" {
		db = db.Where("username LIKE ?", fmt.Sprintf("%%%s%%", req.Search))
	}
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Order("id desc").Find(&ids).Error
	if err != nil {
		return nil, 0, err
	}
	var pos []*user.Enterprise
	err = i.mdb.Model(&user.Enterprise{}).Where("user_id in (?)", ids).Find(&pos).Error
	if err != nil {
		return nil, 0, err
	}
	return pos, total, nil
}
func (i *impl) GetStaffByID(ctx *gin.Context, id int) (*user.User, error) {
	var po *user.User
	err := i.mdb.Model(&user.User{}).Where("id = ?", id).First(&po).Error
	if err != nil {
		return nil, err
	}
	return po, nil
}
func (i *impl) GetEnterpriseByID(ctx *gin.Context, id int) (*user.Enterprise, error) {
	var po *user.Enterprise
	var userInfo user.User
	db := i.mdb.Model(&user.Enterprise{})
	udb := i.mdb.Model(&user.User{}).Select("id,username,role")
	if id == 0 {
		db = db.Where("user_id = ?", utils.GetUserID(ctx))
	} else {
		db = db.Where("id = ?", id)
	}
	err := db.First(&po).Error
	if err != nil {
		return nil, err
	}
	if id == 0 {
		udb = udb.Where("id = ?", utils.GetUserID(ctx))
	} else {
		udb = udb.Where("id = ?", po.UserID)
	}
	err = udb.First(&userInfo).Error
	if err != nil {
		return nil, err
	}
	var name string
	err = i.mdb.Model(&place.School{}).
		Select("name").
		Where("id = ?", po.SchoolID).
		First(&name).Error
	po.SchoolName = name
	po.Username = userInfo.Username
	po.Role = userInfo.Role
	return po, nil
}
func (i *impl) GetEnterpriseByUserID(ctx *gin.Context, userID int) (*user.Enterprise, error) {
	var po *user.Enterprise
	err := i.mdb.Model(&user.Enterprise{}).Where("user_id = ?", userID).First(&po).Error
	if err != nil {
		return nil, err
	}
	return po, nil
}
func (i *impl) GetStudentByID(ctx *gin.Context, id int) (*user.Student, error) {
	var po *user.Student
	var userInfo user.User
	db := i.mdb.Model(&user.Student{})
	udb := i.mdb.Model(&user.User{}).Select("id,username,role")
	if id == 0 {
		db = db.Where("user_id = ?", utils.GetUserID(ctx))
	} else {
		db = db.Where("id = ?", id)
	}
	err := db.First(&po).Error
	if err != nil {
		return nil, err
	}
	if id == 0 {
		udb = udb.Where("id = ?", utils.GetUserID(ctx))
	} else {
		udb = udb.Where("id = ?", po.UserID)
	}
	err = udb.First(&userInfo).Error
	if err != nil {
		return nil, err
	}
	var name string
	err = i.mdb.Model(&place.School{}).
		Select("name").
		Where("id = ?", po.SchoolID).
		First(&name).Error
	po.SchoolName = name
	po.Username = userInfo.Username
	po.Role = userInfo.Role
	return po, nil
}

func (i *impl) UpdateStudent(ctx *gin.Context, student *user.Student) (*user.Student, error) {
	err := i.mdb.Model(&user.Student{}).
		Where("id = ?", student.ID).
		Updates(&student).Error
	if err != nil {
		return nil, err
	}
	return student, nil
}
func (i *impl) UpdateEnterprise(ctx *gin.Context, enterprise *user.Enterprise) (*user.Enterprise, error) {
	err := i.mdb.Model(&user.Enterprise{}).
		Where("id = ?", enterprise.ID).
		Updates(&enterprise).Error
	if err != nil {
		return nil, err
	}
	err = i.mdb.Model(&user.Enterprise{}).
		Where("id = ?", enterprise.ID).
		Updates(map[string]interface{}{
			"updated_at": time.Now().UnixMilli(),
		}).Error
	if err != nil {
		return nil, err
	}
	return enterprise, nil
}
func (i *impl) UpdatePassword(ctx *gin.Context, id int, nowPassword, oldPassword string) error {
	var po user.User
	err := i.mdb.Model(&user.User{}).
		Where("id = ?", id).
		First(&po).Error
	if err != nil {
		return err
	}
	if !utils.BcryptCheck(oldPassword, po.Password) {
		return errors.New("原密码错误")
	}
	err = i.mdb.Model(&user.User{}).
		Where("id = ?", id).
		Update("password", utils.BcryptHash(nowPassword)).Error
	if err != nil {
		return err
	}
	return nil
}
func (i *impl) DeleteStudent(ctx *gin.Context, id int) error {
	var student user.Student
	err := i.mdb.Model(&user.Student{}).Where("id = ?", id).First(&student).Error
	if err != nil {
		return err
	}
	err = i.mdb.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&user.User{}, "id = ?", student.UserID).Error; err != nil {
			return err
		}
		if err := tx.Delete(&user.Student{}, "id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}
	return nil
}
func (i *impl) DeleteStaff(ctx *gin.Context, id int) error {
	err := i.mdb.Delete(&user.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
func (i *impl) DeleteEnterprise(ctx *gin.Context, id int) error {
	var enterprise user.Enterprise
	err := i.mdb.Model(&user.Enterprise{}).Where("id = ?", id).First(&enterprise).Error
	if err != nil {
		return err
	}
	err = i.mdb.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&user.User{}, "id = ?", enterprise.UserID).Error; err != nil {
			return err
		}
		if err := tx.Delete(&user.Enterprise{}, "id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}
	return nil
}
