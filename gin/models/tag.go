package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/astaxie/beego/validation"
	"github.com/Unknwon/com"

	"gin-blog/pkg/e"
	"gin-blog/models"
	"gin-blog/pkg/util"
	"gin-blog/pkg/setting"
)

type Tag struct {
	Model
	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

func GetTagTotal(maps interface {}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}
