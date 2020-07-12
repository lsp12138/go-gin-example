package util

import (
	"github.com/gin-gonic/gin"
	"github.com/lsp12138/go-gin-example/pkg/setting"
	"github.com/unknwon/com"
)

// 工具包，使用github.com/unknwon/com包做类型转换

// GetPage 获取分页页码
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}
