package models

import (
    "time"
)

// 标签的数据库操作

type Tag struct {
    Model

    Name       string `json:"name"`
    CreatedBy  string `json:"created_by"`
    ModifiedBy string `json:"modified_by"`
    State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
    // db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
    tmpDb := db
    if v, ok := maps.(map[string]interface{})["name"]; ok {
        // tmpDb = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", v))
        // 如果模糊查询的值里头有%和_字符，需要用/转义表示
        tmpDb = tmpDb.Where("name LIKE ?", "%"+v.(string)+"%")
    }
    if v, ok := maps.(map[string]interface{})["state"]; ok {
        tmpDb = tmpDb.Where("state = ?", v)
    }
    tmpDb.Where("deleted_on = ?", 0).Offset(pageNum).Limit(pageSize).Find(&tags)
    return
}

func GetTagTotal(maps interface{}) (count int) {
    tmpDb := db
    if v, ok := maps.(map[string]interface{})["name"]; ok {
        tmpDb = tmpDb.Where("name LIKE ?", "%"+v.(string)+"%")
    }
    if v, ok := maps.(map[string]interface{})["state"]; ok {
        tmpDb = tmpDb.Where("state = ?", v)
    }
    tmpDb.Model(&Tag{}).Where("deleted_on = ?", 0).Count(&count)
    return
}

func ExistTagByName(name string) bool {
    var tag Tag
    db.Select("id").Where("name = ?", name).First(&tag)
    if tag.ID > 0 {
        return true
    }
    return false
}

func AddTag(name string, state int, createdBy string) bool {
    db.Create(&Tag{
        Name:      name,
        State:     state,
        CreatedBy: createdBy,
    })
    return true
}

func ExistTagByID(id int) bool {
    var tag Tag
    db.Select("id").Where("id = ?", id).First(&tag)
    if tag.ID > 0 {
        return true
    }
    return false
}

func DeleteTag(id int) bool {
    db.Model(&Tag{}).Where("id = ?", id).Update("deleted_on", time.Now().Unix())
    // db.Where("id = ?", id).Delete(&Tag{})
    return true
}

func EditTag(id int, data interface{}) bool {
    db.Model(&Tag{}).Where("id = ?", id).Updates(data)
    return true
}
