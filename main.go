package main

import (
    "fmt"
    "github.com/lsp12138/go-gin-example/conf"
    "github.com/lsp12138/go-gin-example/routers"
    "net/http"
)

func main() {
    router := routers.InitRouter()
    s := &http.Server{
        Addr:           fmt.Sprintf(":%s", conf.AppConfig.HttpPort),
        Handler:        router,
        ReadTimeout:    60,
        WriteTimeout:   60,
        MaxHeaderBytes: 1 << 20,
    }
    s.ListenAndServe()
}
