package model

import "github.com/yyzcoder/m2go/util"

type Table struct {
	Name    string   `gorm:"column:TABLE_NAME"`
	Comment string   `gorm:"column:TABLE_COMMENT"`
	Columns []Column `gorm:"-"`
}

func (Table) TableName() string {
	return "TABLES"
}

func (t Table) StructName(prefix string) string {
	return util.Snake2Camel(util.TrimPrefix(t.Name, prefix))
}
