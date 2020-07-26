package routers

import (
    "github.com/gin-gonic/gin"
    "github.com/lsp12138/go-gin-example/conf"
    "github.com/lsp12138/go-gin-example/routers/api"
    v1 "github.com/lsp12138/go-gin-example/routers/api/v1"
)

// 注册路由

func InitRouter() *gin.Engine {
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    gin.SetMode(conf.AppConfig.RunMode)

    // 获取token的接口
    r.GET("/auth", api.GetAuth)

    apiv1 := r.Group("/api/v1")
    // 接入jwt中间件到这个路由组
    // apiv1.Use(jwt.JWT())
    {
        // 获取标签列表
        apiv1.GET("/tags", v1.GetTags)
        // 新建标签
        apiv1.POST("/tags", v1.AddTag)
        // 更新指定标签
        apiv1.PUT("/tags/:id", v1.EditTag)
        // 删除指定标签
        apiv1.DELETE("/tags/:id", v1.DeleteTag)

        // 获取文章列表
        apiv1.GET("/articles", v1.GetArticles)
        // 获取指定文章
        apiv1.GET("/articles/:id", v1.GetArticle)
        // 新建文章
        apiv1.POST("/articles", v1.AddArticle)
        // 更新指定文章
        apiv1.PUT("/articles/:id", v1.EditArticle)
        // 删除指定文章
        apiv1.DELETE("/articles/:id", v1.DeleteArticle)
    }
    return r
}
