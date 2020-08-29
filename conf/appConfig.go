package conf

import (
    "github.com/spf13/viper"
)

type appConfig struct {
    RunMode         string
    PageSize        int
    JwtSecret       string
    HttpPort        string
    RuntimeRootPath string
    ImagePrefixUrl  string
    ImageSavePath   string
    ImageMaxSize    int
    ImageAllowExts  []string
}

func loadAppConfig() *appConfig {
    return &appConfig{
        RunMode:         viper.GetString("APP.RUN_MODE"),
        PageSize:        viper.GetInt("APP.PAGE_SIZE"),
        JwtSecret:       viper.GetString("APP.JWT_SECRET"),
        HttpPort:        viper.GetString("APP.HTTP_PORT"),
        RuntimeRootPath: viper.GetString("APP.RUNTIME_ROOT_PATH"),
        ImagePrefixUrl:  viper.GetString("APP.IMAGE_PREFIX_URL"),
        ImageSavePath:   viper.GetString("APP.IMAGE_SAVE_PATH"),
        ImageMaxSize:    viper.GetInt("APP.IMAGE_MAX_SIZE"),
        ImageAllowExts:  viper.GetStringSlice("APP.IMAGE_ALLOW_EXTS"),
    }
}
