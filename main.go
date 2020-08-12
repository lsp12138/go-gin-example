package main

import (
    "fmt"
    "github.com/lsp12138/go-gin-example/conf"
    "github.com/lsp12138/go-gin-example/models"
    "github.com/lsp12138/go-gin-example/routers"
)

func main() {
    router := routers.InitRouter()
    router.Run(fmt.Sprintf(":%s", conf.AppConfig.HttpPort))
    defer models.CloseDB()
}
