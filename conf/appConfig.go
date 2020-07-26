package conf

import "github.com/spf13/viper"

type appConfig struct {
    RunMode   string
    PageSize  int
    JwtSecret string
    HttpPort  string
}

func loadAppConfig() *appConfig {
    return &appConfig{
        RunMode:   viper.GetString("APP.RUN_MODE"),
        PageSize:  viper.GetInt("APP.PAGE_SIZE"),
        JwtSecret: viper.GetString("APP.JWT_SECRET"),
        HttpPort:  viper.GetString("APP.HTTP_PORT"),
    }
}
