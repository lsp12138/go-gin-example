package v1

import (
    "github.com/astaxie/beego/validation"
    "github.com/gin-gonic/gin"
    "github.com/lsp12138/go-gin-example/models"
    "github.com/lsp12138/go-gin-example/pkg/e"
    "github.com/lsp12138/go-gin-example/pkg/util"
    "github.com/unknwon/com"
    "log"
    "net/http"
)

// @Summary 根据id获取文章
// @Tags article文章管理
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles/{id} [get]
func GetArticle(c *gin.Context) {
    id := com.StrTo(c.Param("id")).MustInt()
    valid := validation.Validation{}
    valid.Min(id, 1, "id").Message("ID必须大于0")
    code := e.INVALID_PARAMS
    var data interface{}
    if !valid.HasErrors() {
        if models.ExistArticleByID(id) {
            data = models.GetArticle(id)
            code = e.SUCCESS
        } else {
            code = e.ERROR_NOT_EXIST_ARTICLE
        }
    } else {
        for _, err := range valid.Errors {
            log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
        }
    }
    c.JSON(http.StatusOK, gin.H{
        "code": code,
        "msg":  e.GetMsg(code),
        "data": data,
    })
}

// @Summary 获取文章
// @Tags article文章管理
// @Produce  json
// @Param state query int false "状态"
// @Param tag_id query int false "标签id"
// @Param page_num query int false "分页当前页"
// @Param page_size query int false "分页大小"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles [get]
func GetArticles(c *gin.Context) {
    data := make(map[string]interface{})
    maps := make(map[string]interface{})
    valid := validation.Validation{}

    var state int = -1
    if arg := c.Query("state"); arg != "" {
        state = com.StrTo(arg).MustInt()
        maps["state"] = state

        valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
    }

    var tagId int = -1
    if arg := c.Query("tag_id"); arg != "" {
        tagId = com.StrTo(arg).MustInt()
        maps["tag_id"] = tagId

        valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
    }

    pageSize := util.GetPageSize(c.Query("page_size"))
    pageNum := util.GetPageNum(c.Query("page_num"), pageSize)

    code := e.INVALID_PARAMS
    if !valid.HasErrors() {
        code = e.SUCCESS

        data["lists"] = models.GetArticles(pageNum, pageSize, maps)
        data["total"] = models.GetArticleTotal(maps)

    } else {
        for _, err := range valid.Errors {
            log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "code": code,
        "msg":  e.GetMsg(code),
        "data": data,
    })
}

// @Summary 新增文章
// @Tags article文章管理
// @Produce  json
// @Param title query string true "标题"
// @Param desc query string true "描述"
// @Param content query string true "内容"
// @Param cover_image_url query string true "封面图片"
// @Param created_by query string true "创建者"
// @Param tag_id query int true "标签id"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles [post]
func AddArticle(c *gin.Context) {
    tagId := com.StrTo(c.Query("tag_id")).MustInt()
    title := c.Query("title")
    desc := c.Query("desc")
    content := c.Query("content")
    coverImageUrl := c.Query("cover_image_url")
    createdBy := c.Query("created_by")
    state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()

    valid := validation.Validation{}
    valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
    valid.Required(title, "title").Message("标题不能为空")
    valid.Required(desc, "desc").Message("简述不能为空")
    valid.Required(content, "content").Message("内容不能为空")
    valid.Required(coverImageUrl, "cover_image_url").Message("封面图片不能为空")
    valid.Required(createdBy, "created_by").Message("创建人不能为空")
    valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

    code := e.INVALID_PARAMS
    if !valid.HasErrors() {
        if models.ExistTagByID(tagId) {
            data := make(map[string]interface{})
            data["tag_id"] = tagId
            data["title"] = title
            data["desc"] = desc
            data["content"] = content
            data["cover_image_url"] = coverImageUrl
            data["created_by"] = createdBy
            data["state"] = state

            models.AddArticle(data)
            code = e.SUCCESS
        } else {
            code = e.ERROR_NOT_EXIST_TAG
        }
    } else {
        for _, err := range valid.Errors {
            log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "code": code,
        "msg":  e.GetMsg(code),
        "data": make(map[string]interface{}),
    })
}

// @Summary 修改文章
// @Tags article文章管理
// @Produce  json
// @Param id path int true "id"
// @Param title query string true "标题"
// @Param desc query string true "描述"
// @Param content query string true "内容"
// @Param modified_by query string true "修改者"
// @Param tag_id query int true "标签id"
// @Param state query int true "标签状态"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles/{id} [put]
func EditArticle(c *gin.Context) {
    valid := validation.Validation{}

    id := com.StrTo(c.Param("id")).MustInt()
    tagId := com.StrTo(c.Query("tag_id")).MustInt()
    title := c.Query("title")
    desc := c.Query("desc")
    content := c.Query("content")
    coverImageUrl := c.Query("cover_image_url")
    modifiedBy := c.Query("modified_by")

    var state int = -1
    if arg := c.Query("state"); arg != "" {
        state = com.StrTo(arg).MustInt()
        valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
    }

    valid.Min(id, 1, "id").Message("ID必须大于0")
    valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
    valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
    valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
    valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
    valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

    code := e.INVALID_PARAMS
    if !valid.HasErrors() {
        if models.ExistArticleByID(id) {
            if models.ExistTagByID(tagId) {
                data := make(map[string]interface{})
                if tagId > 0 {
                    data["tag_id"] = tagId
                }
                if title != "" {
                    data["title"] = title
                }
                if desc != "" {
                    data["desc"] = desc
                }
                if content != "" {
                    data["content"] = content
                }
                if coverImageUrl != "" {
                    data["cover_image_url"] = coverImageUrl
                }

                data["modified_by"] = modifiedBy

                models.EditArticle(id, data)
                code = e.SUCCESS
            } else {
                code = e.ERROR_NOT_EXIST_TAG
            }
        } else {
            code = e.ERROR_NOT_EXIST_ARTICLE
        }
    } else {
        for _, err := range valid.Errors {
            log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "code": code,
        "msg":  e.GetMsg(code),
        "data": make(map[string]string),
    })
}

// @Summary 删除文章
// @Tags article文章管理
// @Produce  json
// @Param id path int true "id"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles/{id} [delete]
func DeleteArticle(c *gin.Context) {
    id := com.StrTo(c.Param("id")).MustInt()

    valid := validation.Validation{}
    valid.Min(id, 1, "id").Message("ID必须大于0")

    code := e.INVALID_PARAMS
    if !valid.HasErrors() {
        if models.ExistArticleByID(id) {
            models.DeleteArticle(id)
            code = e.SUCCESS
        } else {
            code = e.ERROR_NOT_EXIST_ARTICLE
        }
    } else {
        for _, err := range valid.Errors {
            log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "code": code,
        "msg":  e.GetMsg(code),
        "data": make(map[string]string),
    })
}
