package impl

import (
	"fmt"
	"gitee.com/xygfm/authorization/apps/admin"
	utils "gitee.com/xygfm/authorization/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"time"
)

func (i *impl) GetUser(ctx *gin.Context, req *admin.LoginRequest) (*admin.User, error) {
	fmt.Println("我是登录业务层")
	var user admin.User
	err := i.mdb.Model(&admin.User{}).
		Where("username", req.Username).
		Where("role", req.Role).
		First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (i *impl) Register(ctx *gin.Context, req *admin.LoginRequest) (string, error) {
	err := i.mdb.Model(&admin.LoginRequest{}).Create(&admin.LoginRequest{
		Username: req.Username,
		Password: req.Password,
		Role:     req.Role,
	}).Error
	if err != nil {
		fmt.Println("err", err)
		return "", err
	}
	return "", nil
}

func (i *impl) ListArea(ctx *gin.Context, id string) ([]*admin.Areas, error) {
	var po []*admin.Areas
	db := i.mdb.Model(&admin.Areas{})
	if id != "0" {
		db.Where("parent_id = ?", id)
	}
	if utils.GetUserID(ctx) != 1 {
		db.Where("state = ?", 1)
	}
	err := db.Find(&po).Error
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	fmt.Println("po", po)
	// 使用 Map 缓存所有菜单
	areaMap := make(map[int]*admin.Areas)
	for _, item := range po {
		areaMap[item.ID] = item
	}

	// 构建树形结构
	var tree []*admin.Areas
	for _, item := range po {
		if item.ParentID == 1 {
			// 根节点
			tree = append(tree, item)
		} else {
			// 尝试找父节点
			parent, ok := areaMap[item.ParentID]
			if ok {
				// 找到父节点，添加为子节点
				parent.Children = append(parent.Children, item)
			} else {
				// 找不到父节点，作为顶层节点处理
				tree = append(tree, item)
			}
		}
	}

	return tree, nil
}

func (i *impl) GetStatistics(ctx *gin.Context) (*admin.Statistics, error) {
	var po admin.Statistics
	//err := i.mdb.Model(&admin.Statistics{}).First(&po).Error
	//if err != nil {
	//	return nil, err
	//}
	var roleCounts []admin.RoleCount
	i.mdb.Model(&admin.User{}).
		Select("role, COUNT(*) as count").
		Group("role").
		Find(&roleCounts)
	// 遍历结果并赋值
	for _, rc := range roleCounts {
		if rc.Role == 1 {
			po.StudentSum = int(rc.Count)
		} else if rc.Role == 2 {
			po.EnterpriseSum = int(rc.Count)
		}
	}

	i.mdb.Model(&admin.User{}).
		Select("role, COUNT(*) as count").
		Where("created_at >= ?", time.Now().AddDate(0, 0, -7).UnixMilli()).
		Group("role").
		Find(&roleCounts)
	for _, rc := range roleCounts {
		if rc.Role == 1 {
			po.NewStudentSum = int(rc.Count)
		} else if rc.Role == 2 {
			po.NewEnterpriseSum = int(rc.Count)
		}
	}
	jobData, _ := i.getDailyJobCounts()
	studentData, _ := i.getDailyUserCounts(1)
	enterpriseData, _ := i.getDailyUserCounts(2)
	for _, item := range jobData {
		po.NewJob = append(po.NewJob, cast.ToInt(item.Count))
		po.Date = append(po.Date, item.Date)
	}
	for _, item := range studentData {
		po.NewStudent = append(po.NewStudent, cast.ToInt(item.Count))
	}
	for _, item := range enterpriseData {
		po.NewEnterprise = append(po.NewEnterprise, cast.ToInt(item.Count))
	}
	//po.NewStudent = []int{5, 7, 10, 45, 55, 23, 67}
	//po.NewEnterprise = []int{6, 10, 22, 43, 9, 25, 32}
	//po.NewJob = []int{66, 75, 80, 34, 56, 67, 34}
	//po.Date = utils.GetRecentSevenDays()
	return &po, nil
}

func (i *impl) SetUserUpdateAt(ctx *gin.Context, id int) error {
	err := i.mdb.Model(&admin.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"update_at": time.Now().UnixMilli(),
		}).Error
	if err != nil {
		return err
	}
	return nil
}
func (i *impl) getDailyRegisterCounts() ([]admin.RegisterCount, error) {
	var results []admin.RegisterCount

	// 获取当前时间
	now := time.Now()

	// 计算7天前的时间
	sevenDaysAgo := now.AddDate(0, 0, -7)

	// 查询每一天的注册人数
	if err := i.mdb.Table("users").
		Select("FROM_UNIXTIME(created_at / 1000, '%Y-%m-%d') AS date, COUNT(*) AS count").
		Where("created_at >= ?", sevenDaysAgo.UnixNano()/int64(time.Millisecond)).
		Group("FROM_UNIXTIME(created_at / 1000, '%Y-%m-%d')").
		Order("date ASC").
		Find(&results).Error; err != nil {
		return nil, err
	}

	// 确保结果包含所有7天的数据，即使某些天没有注册用户
	allDates := make([]string, 7)
	for i := 0; i < 7; i++ {
		day := sevenDaysAgo.AddDate(0, 0, i)
		allDates[i] = day.Format("2006-01-02")
	}

	// 补全缺失日期的注册人数为0
	completeResults := make([]admin.RegisterCount, 7)
	for i, date := range allDates {
		found := false
		for _, result := range results {
			if result.Date == date {
				completeResults[i] = result
				found = true
				break
			}
		}
		if !found {
			completeResults[i] = admin.RegisterCount{Date: date, Count: 0}
		}
	}
	fmt.Println("completeResults", completeResults)

	return completeResults, nil
}

func (i *impl) getDailyJobCounts() ([]admin.RegisterCount, error) {
	var results []admin.RegisterCount
	// 获取当前时间
	now := time.Now()
	// 计算7天前的时间
	sevenDaysAgo := now.AddDate(0, 0, -7)
	// 查询每一天的注册人数
	if err := i.mdb.Table("jobs").
		Select("FROM_UNIXTIME(updated_at / 1000, '%Y-%m-%d') AS date, COUNT(*) AS count").
		Where("updated_at >= ?", sevenDaysAgo.UnixNano()/int64(time.Millisecond)).
		Group("FROM_UNIXTIME(updated_at / 1000, '%Y-%m-%d')").
		Order("date ASC").
		Find(&results).Error; err != nil {
		return nil, err
	}

	// 确保结果包含所有7天的数据，即使某些天没有注册用户
	allDates := make([]string, 7)
	for i := 0; i < 7; i++ {
		day := sevenDaysAgo.AddDate(0, 0, i)
		allDates[i] = day.Format("2006-01-02")
	}
	// 补全缺失日期的注册人数为0
	completeResults := make([]admin.RegisterCount, 7)
	for i, date := range allDates {
		found := false
		for _, result := range results {
			if result.Date == date {
				completeResults[i] = result
				found = true
				break
			}
		}
		if !found {
			completeResults[i] = admin.RegisterCount{Date: date, Count: 0}
		}
	}
	fmt.Println("completeResults", completeResults)
	return completeResults, nil
}
func (i *impl) getDailyUserCounts(role int) ([]admin.RegisterCount, error) {
	var results []admin.RegisterCount
	// 获取当前时间
	now := time.Now()
	// 计算7天前的时间
	sevenDaysAgo := now.AddDate(0, 0, -7)
	// 查询每一天的注册人数
	if err := i.mdb.Table("users").
		Select("FROM_UNIXTIME(updated_at / 1000, '%Y-%m-%d') AS date, COUNT(*) AS count").
		Where("role = ?", role).
		Where("updated_at >= ?", sevenDaysAgo.UnixNano()/int64(time.Millisecond)).
		Group("FROM_UNIXTIME(updated_at / 1000, '%Y-%m-%d')").
		Order("date ASC").
		Find(&results).Error; err != nil {
		return nil, err
	}

	// 确保结果包含所有7天的数据，即使某些天没有用户
	allDates := make([]string, 7)
	for i := 0; i < 7; i++ {
		day := sevenDaysAgo.AddDate(0, 0, i)
		allDates[i] = day.Format("2006-01-02")
	}
	// 补全缺失日期的注册人数为0
	completeResults := make([]admin.RegisterCount, 7)
	for i, date := range allDates {
		found := false
		for _, result := range results {
			if result.Date == date {
				completeResults[i] = result
				found = true
				break
			}
		}
		if !found {
			completeResults[i] = admin.RegisterCount{Date: date, Count: 0}
		}
	}
	fmt.Println("completeResults", completeResults)
	return completeResults, nil
}
