package models

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/lsp12138/go-gin-example/conf"
    "log"
    "time"
)

// models的初始化使用，使用github.com/jinzhu/gorm包操作数据库

var db *gorm.DB

type Model struct {
    ID         int `gorm:"primary_key" json:"id"`
    CreatedOn  int `json:"created_on"`
    ModifiedOn int `json:"modified_on"`
    DeletedOn  int `json:"deleted_on"`
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
    // 注册 Callbacks
    db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
    db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
}

func CloseDB() {
    defer db.Close()
}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
    if !scope.HasError() {
        nowTime := time.Now().Unix()
        if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
            if createTimeField.IsBlank {
                createTimeField.Set(nowTime)
            }
        }
        if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
            if modifyTimeField.IsBlank {
                modifyTimeField.Set(nowTime)
            }
        }
    }
}

// updateTimeStampForUpdateCallback will set `ModifyTime` when updating
// scope.Get(...) 根据入参获取设置了字面值的参数，
// 例如本文中是 gorm:update_column ，它会去查找含这个字面值的字段属性
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
    if _, ok := scope.Get("gorm:update_column"); !ok {
        scope.SetColumn("ModifiedOn", time.Now().Unix())
    }
}
