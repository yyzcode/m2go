package model

import (
	"fmt"
	"github.com/yyzcoder/m2go/util"
	"strings"
)

type Column struct {
	Name       string  `gorm:"column:COLUMN_NAME"`
	DataType   string  `gorm:"column:DATA_TYPE"`
	Comment    string  `gorm:"column:COLUMN_COMMENT"`
	Default    *string `gorm:"column:COLUMN_DEFAULT"`
	ColumnType string  `gorm:"column:COLUMN_TYPE"`
	IsNullable string  `gorm:"column:IS_NULLABLE"`
}

//mysql有符号与go对应
var mt2gt = map[string]string{
	"tinyint":    "int8",
	"smallint":   "int16",
	"int":        "int",
	"integer":    "int",
	"mediumint":  "int",
	"bigint":     "int64",
	"decimal":    "float64",
	"float":      "float64",
	"char":       "string",
	"varchar":    "string",
	"text":       "string",
	"mediumtext": "string",
	"longtext":   "string",
	"date":       "time.Time",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
}

//mysql无符号与go类型对应
var umt2gt = map[string]string{
	"tinyint":    "uint8",
	"smallint":   "uint16",
	"int":        "uint",
	"integer":    "uint",
	"mediumint":  "uint",
	"bigint":     "uint64",
	"decimal":    "float64",
	"float":      "float64",
	"char":       "string",
	"varchar":    "string",
	"text":       "string",
	"mediumtext": "string",
	"longtext":   "string",
	"date":       "time.Time",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
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

func (c Column) FieldDefault() (s string) {
	if c.Default != nil && *c.Default != "" {
		s = fmt.Sprintf("//default:%s", *c.Default)
	}
	return s
}

func (c Column) fieldJson() string {
	return fmt.Sprintf("`json:\"%s\"`", c.Name)
}

func (c Column) FieldType() (s string) {
	m := mt2gt
	if strings.Contains(c.ColumnType, "unsigned") {
		m = umt2gt
	}
	if v, ok := m[c.DataType]; ok {
		if c.IsNullable == "YES" {
			return "*" + v
		}
		return v
	}
	return "any"
}

func (Column) TableName() string {
	return "COLUMNS"
}
