m2go
===============

m2go是一款根据mysql数据表信息生成对应go结构体的工具。

## 约定
根据日常使用习惯，对生成规则做如下约定
* 文件生成的位置为执行命令的当前目录
* 生成的文件名为 表名.go（不含表前缀）
* 生成的结构体名为表名大驼峰规范（不含表前缀）
* 生成的结构体字段为表字段名的大驼峰规范
* 生成的json标签与数据库字段名保持一致
* 如产生同名文件，新文件会添加时间后缀，使用 -w 可以覆盖同名文件

## 安装

~~~
go install github.com/yyzcoder/m2go
~~~

## 快速使用

生成所有表结构体

~~~
m2go -host 127.0.0.1:3306 -u user:password -db dbname
~~~

生成指定表结构体

~~~
m2go -host 127.0.0.1:3306 -u user:password -db dbname table1 table2
~~~

生成带json标签的结构体
~~~
m2go -host 127.0.0.1:3306 -u user:password -db dbname -j
~~~

带前缀的表名，示例中将会匹配 prefix_user 表，注意 prefix_ 中的下划线本身就是前缀的一部分
~~~
m2go -host 127.0.0.1:3306 -u user:password -db dbname -j -p prefix_ user
~~~

生成结果
```go
package test

// Province 省份信息表
type Province struct {
	Id      int    `json:"id"`
	Myid    int    `json:"myid"`     //自定义ID
	AllName string `json:"all_name"` //省份全程 //default:北京市
	Name    string `json:"name"`     //名称缩写
	Pycode  string `json:"pycode"`   //拼音码
	Pinyin  string `json:"pinyin"`   //拼音
}
```

## 参数

* -host 127.0.0.1:3306 数据库地址与端口
* -u root:123456 数据库用户名与密码
* -db dbname 数据库名
* -p table_prefix 数据库前缀
* -w 如果生成同名文件则覆盖已有文件
* -j 生成结构体字段的json标签
* -dv 生成结构体字段在数据库的默认值备注
