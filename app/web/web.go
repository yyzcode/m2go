package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yyzcode/m2go/app/web/static/css"
	"github.com/yyzcode/m2go/app/web/view"
	db2 "github.com/yyzcode/m2go/db"
	"github.com/yyzcode/m2go/model"
	"gorm.io/gorm"
)

type WebOption struct {
	Addr     string
	DbHost   string
	DbUser   string
	Database string
}

var db *gorm.DB
var webOption WebOption

func Run(option WebOption) error {
	webOption = option
	r := gin.Default()
	//数据库连接
	dsn := fmt.Sprintf("%s@tcp(%s)/information_schema?charset=utf8mb4&parseTime=True&loc=Local",
		option.DbUser,
		option.DbHost,
	)
	var err error
	db, err = db2.Connect(dsn)
	if err != nil {
		return err
	}

	r.GET("/tables", tableList)

	r.GET("/static/css/reset.css", func(context *gin.Context) {
		context.Header("content-type", "text/css")
		context.String(200, css.Reset)
	})
	
	err = r.Run(option.Addr)
	if err != nil {
		return err
	}
	return nil
}

func tableList(ctx *gin.Context) {
	var err error
	tables := make([]model.Table, 0, 0)
	query := db.Where("TABLE_SCHEMA = ?", webOption.Database)
	err = query.Find(&tables).Error

	render, err := view.Build("/model", gin.H{
		"title":  "m2go - 模型生成器",
		"tables": tables,
	})

	if err != nil {
		fmt.Println("render html err:", err)
	}
	if err = render.Render(ctx.Writer); err != nil {
		fmt.Println("render html err:", err)
	}
}
