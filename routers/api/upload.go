package api

import (
    "github.com/gin-gonic/gin"
    "github.com/lsp12138/go-gin-example/pkg/e"
    "github.com/lsp12138/go-gin-example/pkg/upload"
    "log"
    "net/http"
)

// @Summary 上传图片
// @Tags 上传图片
// @Description 上传文章封面图片
// @Produce  json
// @Accept multipart/form-data
// @Param image formData file true "图片"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /upload [put]
func UploadImage(c *gin.Context) {
    code := e.SUCCESS
    data := make(map[string]string)

    file, image, err := c.Request.FormFile("image")
    if err != nil {
        log.Println(err)
        code = e.ERROR
        c.JSON(http.StatusOK, gin.H{
            "code": code,
            "msg":  e.GetMsg(code),
            "data": data,
        })
    }
    if image == nil {
        code = e.INVALID_PARAMS
    } else {
        imageName := upload.GetImageName(image.Filename)
        fullPath := upload.GetImageFullPath()
        savePath := upload.GetImagePath()

        src := fullPath + imageName
        if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
            code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
        } else {
            err := upload.CheckImage(fullPath)
            if err != nil {
                log.Println(err)
                code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
            } else if err := c.SaveUploadedFile(image, src); err != nil {
                log.Println(err)
                code = e.ERROR_UPLOAD_SAVE_IMAGE_FAIL
            } else {
                data["image_url"] = upload.GetImageFullUrl(imageName)
                data["image_save_url"] = savePath + imageName
            }
        }
    }
    c.JSON(http.StatusOK, gin.H{
        "code": code,
        "msg":  e.GetMsg(code),
        "data": data,
    })
}
