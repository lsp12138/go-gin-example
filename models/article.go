package models

import (
    "github.com/jinzhu/gorm"
    "time"
)

// 文章的数据库操作

type Article struct {
    Model

    TagID int `json:"tag_id" gorm:"index"`
    Tag   Tag `json:"tag"`

    Title      string `json:"title"`
    Desc       string `json:"desc"`
    Content    string `json:"content"`
    CreatedBy  string `json:"created_by"`
    ModifiedBy string `json:"modified_by"`
    State      int    `json:"state"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("CreatedOn", time.Now().Unix())
    return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
    scope.SetColumn("ModifiedOn", time.Now().Unix())
    return nil
}

func ExistArticleByID(id int) bool {
    var article Article
    db.Select("id").Where("id = ?", id).First(&article)
    if article.ID > 0 {
        return true
    }
    return false
}

// Article表是如何关联到Tag表的？
// gorm通过Related进行关联查询
func GetArticle(id int) (article Article) {
    db.Where("id = ?", id).First(&article)
    db.Model(&article).Related(&article.Tag)
    return
}

// Preload可以查询出每一项的关联Tag
// 它会执行两条 SQL，分别是 SELECT * FROM blog_articles;
// 和 SELECT * FROM blog_tag WHERE id IN (1,2,3,4);
// 那么在查询出结构后，gorm将其填充到Article的Tag中，避免了循环查询
// 那么有没有别的办法呢，大致是两种
// gorm的Join 和 循环Related
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
    db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
    return
}

func GetArticleTotal(maps interface{}) (count int) {
    db.Model(&Article{}).Where(maps).Count(&count)
    return
}

func EditArticle(id int, data interface{}) bool {
    db.Model(&Article{}).Where("id = ?", id).Updates(data)
    return true
}

func AddArticle(data map[string]interface{}) bool {
    db.Create(&Article{
        TagID:     data["tag_id"].(int),
        Title:     data["title"].(string),
        Desc:      data["desc"].(string),
        Content:   data["content"].(string),
        CreatedBy: data["created_by"].(string),
        State:     data["state"].(int),
    })
    return true
}

func DeleteArticle(id int) bool {
    db.Where("id = ?", id).Delete(&Article{})
    return true
}
