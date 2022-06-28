package command

import (
	"fmt"
	db2 "github.com/yyzcode/m2go/db"
	"github.com/yyzcode/m2go/model"
	"gorm.io/gorm"
)

type CmdOption struct {
	Overwrite    bool
	JsonFlag     bool
	DefaultValue bool
	Prefix       string
	DbHost       string
	DbUser       string
	Database     string
	Tables       []string
}

var db *gorm.DB

func Run(option CmdOption) error {
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
	//拿到要生成结构体的所有表名，如为空则生成所有表的结构体
	tableNames := option.Tables
	if option.Prefix != "" {
		//修正带前缀的表名
		for i := 0; i < len(tableNames); i++ {
			tableNames[i] = option.Prefix + tableNames[i]
		}
	}
	tables := make([]model.Table, 0, 0)
	query := db.Where("TABLE_SCHEMA = ?", option.Database)
	if len(tableNames) > 0 {
		query = query.Where("TABLE_NAME IN ?", tableNames)
	}
	err = query.Find(&tables).Error

	if err != nil {
		return err
	}

	//获取每张表的字段信息
	for i := range tables {
		err = db.
			Where("TABLE_SCHEMA = ?", option.Database).
			Where("TABLE_NAME = ?", tables[i].Name).
			Find(&tables[i].Columns).
			Error
		if err != nil {
			return err
		}

	}
	return genStruct(tables, option)
}

func genStruct(tables []model.Table, option CmdOption) error {
	for _, table := range tables {
		file := model.BuildGoFile(table, model.FOption{
			Prefix:      option.Prefix,
			JsonFlag:    option.JsonFlag,
			DefaultFlag: option.DefaultValue,
		})
		err := file.Output(option.Overwrite)
		if err != nil {
			return err
		}
		fmt.Printf("generate: %s\n", file.Name)
	}
	return nil
}
