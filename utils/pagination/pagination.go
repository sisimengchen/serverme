package pagination

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/sisimengchen/serverme/configs"
)

var (
	defaultPageSize = configs.Viper.GetString("pageSize")
)

// type PaginationResponse struct {
// 	List         []interface{}    `json:"list"`
// 	PageNum      int64            `json:"pageNum"`
// 	PageSize     string           `json:"pageSize"`
// 	TotalCount   string           `json:"totalCount"`
// 	TotalPage    string           `json:"totalPage"`
// }

// func PaginationResource(code int, msg string, data interface{}) (int, interface{}) {
// 	return 200, gin.H{
// 		"code": code,
// 		"msg":  msg,
// 		"data": data,
// 	}
// }

func GetPage(ctx *gin.Context) (int, int) {
	pageSize := ctx.DefaultQuery("pageSize", defaultPageSize)
	pageNum := ctx.DefaultQuery("pageNum", "1")
	// 限制一页返回的数量，如果返回的非法值，则设置为 -1
	limit, err := strconv.Atoi(pageSize)
	if err != nil {
		limit = -1
	}
	offset, err := strconv.Atoi(pageNum)
	if err != nil || limit <= 0 {
		offset = -1
	} else {
		offset = (offset - 1) * limit
	}
	return offset, limit
}