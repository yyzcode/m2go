package model

import (
	"fmt"
	"github.com/yyzcoder/m2go/util"
)

type Column struct {
	Name     string `gorm:"column:COLUMN_NAME"`
	DataType string `gorm:"column:DATA_TYPE"`
	Comment  string `gorm:"column:COLUMN_COMMENT"`
	Default  string `gorm:"column:COLUMN_DEFAULT"`
}

func (c Column) fieldName() string {
	return util.Snake2Camel(c.Name)
}

func (c Column) FieldNote() (s string) {
	if c.Comment != "" {
		s = fmt.Sprintf("//%s", c.Comment)
	}
	return s
}

func (c Column) fieldJson() string {
	return fmt.Sprintf("`json:\"%s\"`", c.Name)
}

func (c Column) FieldType() (s string) {
	switch c.DataType {
	case "tinyint":
		s = "int8"
	case "smallint":
		s = "int16"
	case "int":
		s = "int"
	case "integer":
		s = "int"
	case "mediumint":
		s = "int"
	case "bigint":
		s = "int64"
	case "decimal":
		s = "float64"
	case "float":
		s = "float64"
	case "char":
		s = "string"
	case "varchar":
		s = "string"
	case "text":
		s = "string"
	case "mediumtext":
		s = "string"
	case "longtext":
		s = "string"
	case "binary":
		s = "[]byte"
	case "blob":
		s = "[]byte"
	case "tinyblob":
		s = "[]byte"
	case "mediumblob":
		s = "[]byte"
	case "longblob":
		s = "[]byte"
	case "date":
		//考虑到时间类型的零值无法写入数据库，所以用指针
		s = "*time.Time"
	case "datetime":
		s = "*time.Time"
	case "timestamp":
		s = "*time.Time"
	default:
		s = "any"
	}
	return s
}

func (Column) TableName() string {
	return "COLUMNS"
}
