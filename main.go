package main

import (
    "fmt"
    "github.com/lsp12138/go-gin-example/conf"
    "github.com/lsp12138/go-gin-example/models"
    "github.com/lsp12138/go-gin-example/routers"
)

// @title go-gin-example
// @version 1.0
// @description  "gin框架demo"
// @termsOfService http://github.com
// @contact.name API Support
// @contact.url http://www.cnblogs.com
// @contact.email ×××@qq.com
func main() {
    // 初始化配置
    conf.Setup()
    models.Setup()
    router := routers.InitRouter()
    router.Run(fmt.Sprintf(":%s", conf.AppConfig.HttpPort))
    defer models.CloseDB()
}
