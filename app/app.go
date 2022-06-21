package app

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/yyzcode/m2go/db"
	"github.com/yyzcode/m2go/model"
)

var overwrite bool
var jsonFlag bool
var defaultValue bool
var prefix string

func Run(ctx *cli.Context) error {
	//获取输入信息
	overwrite = ctx.Bool("overwrite")
	defaultValue = ctx.Bool("default_value")
	jsonFlag = ctx.Bool("json")
	prefix = ctx.String("prefix")
	database := ctx.String("database")
	//数据库连接
	dsn := fmt.Sprintf("%s@tcp(%s)/information_schema?charset=utf8mb4&parseTime=True&loc=Local",
		ctx.String("user"),
		ctx.String("mysqlhost"),
	)
	dbConn, err := db.Connect(dsn)
	if err != nil {
		return err
	}
	//拿到要生成结构体的所有表名，如为空则生成所有表的结构体
	tableNames := ctx.Args().Slice()
	if prefix != "" {
		//修正带前缀的表名
		for i := 0; i < len(tableNames); i++ {
			tableNames[i] = prefix + tableNames[i]
		}
	}
	tables := make([]model.Table, 0, 0)
	query := dbConn.Where("TABLE_SCHEMA = ?", database)
	if len(tableNames) > 0 {
		query = query.Where("TABLE_NAME IN ?", tableNames)
	}
	err = query.Find(&tables).Error

	if err != nil {
		return err
	}

	//获取每张表的字段信息
	for i := range tables {
		err = dbConn.
			Where("TABLE_SCHEMA = ?", database).
			Where("TABLE_NAME = ?", tables[i].Name).
			Find(&tables[i].Columns).
			Error
		if err != nil {
			return err
		}

	}
	return genStruct(tables)
}

func genStruct(tables []model.Table) error {
	for _, table := range tables {
		file := model.BuildGoFile(table, model.FOption{
			Prefix:      prefix,
			JsonFlag:    jsonFlag,
			DefaultFlag: defaultValue,
		})
		err := file.Output(overwrite)
		if err != nil {
			return err
		}
		fmt.Printf("generate: %s\n", file.Name)
	}
	return nil
}
