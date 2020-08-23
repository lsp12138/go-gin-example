package upload

import (
    "github.com/lsp12138/go-gin-example/conf"
    "github.com/lsp12138/go-gin-example/pkg/file"
    "github.com/lsp12138/go-gin-example/pkg/util"
    "mime/multipart"
    "path"
    "strings"
)

func GetImageFullUrl(name string) string {
    return conf.AppConfig.ImagePrefixUrl + "/" + GetImagePath() + name
}

func GetImageName(name string) string {
    ext := path.Ext(name)
    fileName := strings.TrimSuffix(name, ext)
    fileName = util.EncodeMD5(fileName)
    return fileName + ext
}

func GetImagePath() string {
    return conf.AppConfig.ImageSavePath
}

func GetImageFullPath() string {
    return conf.AppConfig.RuntimeRootPath + GetImagePath()
}

func CheckImageExt(fileName string) bool {
    ext := file.GetExt(fileName)
    for _, allowExt := range conf.AppConfig.ImageAllowExts {
        if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
            return true
        }
    }
    return false
}

func CheckImageSize(f multipart.File) bool {
}
