package model

import (
	"fmt"
	"github.com/yyzcode/m2go/util"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type GoFile struct {
	Name       string
	PkgName    string   //包名
	IptPkg     []string //引用库
	Comment    string
	StructName string
	Fields     []Column
	Option     FOption
}

type FOption struct {
	JsonFlag    bool
	DefaultFlag bool
	Prefix      string
}

// Output 生成文件
func (g *GoFile) Output(overwrite bool) error {
	var file *os.File
	var err error

	if fileExist(g.Name) && !overwrite {
		//如果不覆盖旧文件，则新文件名后面加上时间戳
		g.Name = strings.Replace(g.Name, ".go", strconv.Itoa(int(time.Now().Unix()))+".go", 1)
	}

	file, err = os.OpenFile(g.Name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer func() {
		cmd := exec.Command("gofmt", "-w", g.Name)
		if err = cmd.Run(); err != nil {
			fmt.Println("warning: format code style fail ", err)
		}
	}()
	defer file.Close()

	//写入包名
	file.Write([]byte(fmt.Sprintf("package %s\n\n", g.PkgName)))
	//写入import
	if len(g.IptPkg) > 0 {
		for _, p := range g.IptPkg {
			file.Write([]byte(fmt.Sprintf("import \"%s\"\n\n", p)))
		}
	}
	if g.Comment != "" {
		//写入结构体注释
		file.Write([]byte(fmt.Sprintf("// %s %s\n", g.StructName, g.Comment)))
	}
	//写入结构体名
	file.Write([]byte(fmt.Sprintf("type %s struct {\n", g.StructName)))
	//写入结构体字段
	for _, field := range g.Fields {
		file.Write([]byte("\t"))
		file.Write([]byte(field.fieldName()))
		file.Write([]byte(" "))
		file.Write([]byte(field.FieldType()))
		file.Write([]byte(" "))
		if g.Option.JsonFlag {
			file.Write([]byte(field.fieldJson()))
			file.Write([]byte(" "))
		}
		file.Write([]byte(field.FieldNote()))
		file.Write([]byte(" "))
		if g.Option.DefaultFlag {
			file.Write([]byte(field.FieldDefault()))
			file.Write([]byte(" "))
		}
		file.Write([]byte("\n"))
	}
	file.Write([]byte("}"))
	return nil
}

// BuildGoFile 将table信息转化为更适合用来描述go文件构成的结构体
func BuildGoFile(table Table, option FOption) GoFile {
	file := GoFile{
		Option: option,
	}
	//确定go文件名
	file.Name = fmt.Sprintf("%s.go", util.TrimPrefix(table.Name, option.Prefix))

	//使用当前文件夹作为包名
	currentPath, _ := os.Getwd()
	separator := string([]byte{os.PathSeparator})
	dirs := strings.Split(currentPath, separator)
	packageName := dirs[len(dirs)-1]
	file.PkgName = strings.ReplaceAll(packageName, "-", "_")

	//其他信息
	file.Comment = table.Comment
	file.StructName = table.StructName(option.Prefix)
	file.Fields = table.Columns

	flag := false //是否有time包引入
	for _, f := range file.Fields {
		if f.FieldType() == "*time.Time" || f.FieldType() == "time.Time" {
			flag = true
		}
	}
	//如果有time包依赖，把time写进iptPkg
	if flag {
		file.IptPkg = append(file.IptPkg, "time")
	}
	return file
}

func fileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
