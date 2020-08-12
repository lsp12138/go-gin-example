package conf

import "github.com/spf13/viper"

type dbConfig struct {
    Type        string
    User        string
    Password    string
    Host        string
    Name        string
    TablePrefix string
}

func loadDbConfig() *dbConfig {
    return &dbConfig{
        Type:        viper.GetString("DB.TYPE"),
        User:        viper.GetString("DB.USER"),
        Password:    viper.GetString("DB.PASSWORD"),
        Host:        viper.GetString("DB.HOST"),
        Name:        viper.GetString("DB.NAME"),
        TablePrefix: viper.GetString("DB.TABLE_PREFIX"),
    }
}
