package pagination

import (
	"fmt"
	"strconv"
	"github.com/kataras/iris"
)

const DefaultPageSize = "20"

func GetPage(ctx iris.Context) (int, int) {
	pageSize := ctx.URLParamDefault("pageSize", DefaultPageSize)
	pageNum := ctx.URLParamDefault("pageNum", "1")
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
	fmt.Printf("GetPage:%d,%d", offset, limit)
	return offset, limit
}