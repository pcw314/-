package impl

import (
	"gitee.com/xygfm/authorization/apps/position"
	"github.com/gin-gonic/gin"
)

func (i *impl) ListCollect(ctx *gin.Context, userID int) ([]*position.Collect, error) {
	i.mdb.Model(&position.Collect{}).Where("user_id = ?", userID)
	return nil, nil
}
func (i *impl) CreateCollect(ctx *gin.Context, collect position.Collect) (*position.Collect, error) {
	return nil, nil
}
func (i *impl) DeleteCollect(ctx *gin.Context, id int) error {
	return nil
}
