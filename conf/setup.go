package conf

import (
    "github.com/fsnotify/fsnotify"
    "github.com/spf13/viper"
    "log"
)

const (
    configFilePath = "./conf/config.yaml"
    configFileType = "yaml"
)

var (
    AppConfig *appConfig
    DBConfig  *dbConfig
)

func Setup() {
    loadConfigFromYaml()
    AppConfig = loadAppConfig()
    DBConfig = loadDbConfig()
    watchConfig()
}

func loadConfigFromYaml() {
    viper.SetConfigFile(configFilePath)
    viper.SetConfigType(configFileType)
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("read config failed: %v", err)
    }
}

// 监控配置文件变化实现热更新
func watchConfig() {
    viper.WatchConfig()
    // 如果文件配置文件发生改变就会触发这个函数
    viper.OnConfigChange(func(ev fsnotify.Event) {
        log.Printf("Config file changed: %s", ev.Name)
    })
}
