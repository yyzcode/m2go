package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yyzcode/m2go/app/web/view"
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

		r, err := view.Build("/model", gin.H{"title": "m2go - 模型生成器"})
		if err != nil {
			fmt.Println("render html err:", err)
		}
		if err = r.Render(ctx.Writer); err != nil {
			fmt.Println("render html err:", err)
		}
	})

	err := r.Run(option.Addr)
	if err != nil {
		return err
	}
	return nil
}
