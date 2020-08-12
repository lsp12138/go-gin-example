package util

import (
    "github.com/unknwon/com"
)

// 工具包，使用github.com/unknwon/com包做类型转换

// GetPageNum 获取分页页码
// 设置page_num=1为第一页
func GetPageNum(pageNum string, pageSize int) int {
    result := 0
    page, _ := com.StrTo(pageNum).Int()
    if page > 0 {
        result = (page - 1) * pageSize
    }
    return result
}

// GetPageSize 获取分页大小
// 默认大小为10
func GetPageSize(pageSize string) int {
    result := 10
    size, _ := com.StrTo(pageSize).Int()
    if size > 0 {
        result = size
    }
    return result
}
