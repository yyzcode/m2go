package web

import (
	"github.com/gin-gonic/gin"
)

type WebOption struct {
	Addr     string
	DbHost   string
	DbUser   string
	Database string
}

func Run(option WebOption) error {
	r := gin.Default()

	r.GET("/model", func(ctx *gin.Context) {
		ctx.Header("content-type", "text/html")
		ctx.String(200, html)
	})

	err := r.Run(option.Addr)
	if err != nil {
		return err
	}
	return nil
}

var html = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>m2go - 模型生成器</title>
</head>
<body>
功能开发中，敬请期待
</body>
</html>
`
