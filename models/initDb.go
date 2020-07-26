package models

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/lsp12138/go-gin-example/conf"
    "log"
)

// models的初始化使用，使用github.com/jinzhu/gorm包操作数据库

var db *gorm.DB

type Model struct {
    ID         int `gorm:"primary_key" json:"id"`
    CreatedOn  int `json:"created_on"`
    ModifiedOn int `json:"modified_on"`
}

func init() {
    var err error
    // 这里使用 := 时会创建一个新的db变量把全局变量db覆盖掉，导致空指针报错
    db, err = gorm.Open(conf.DBConfig.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
        conf.DBConfig.User,
        conf.DBConfig.Password,
        conf.DBConfig.Host,
        conf.DBConfig.Name))
    if err != nil {
        log.Println(err)
    }
    gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
        return conf.DBConfig.TablePrefix + defaultTableName
    }
    db.SingularTable(true)
    db.LogMode(true)
    db.DB().SetMaxIdleConns(10)
    db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
    defer db.Close()
}
