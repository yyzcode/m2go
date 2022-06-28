package app

import (
	"github.com/urfave/cli/v2"
	"github.com/yyzcode/m2go/app/command"
	"github.com/yyzcode/m2go/app/web"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Run(ctx *cli.Context) error {

	if ctx.String("web") == "" {
		return command.Run(command.CmdOption{
			Overwrite:    ctx.Bool("overwrite"),
			JsonFlag:     ctx.Bool("json"),
			DefaultValue: ctx.Bool("default_value"),
			Prefix:       ctx.String("prefix"),
			DbUser:       ctx.String("user"),
			DbHost:       ctx.String("mysqlhost"),
			Database:     ctx.String("database"),
			Tables:       ctx.Args().Slice(),
		})
	}

	return web.Run(web.WebOption{
		Addr:     ctx.String("web"),
		DbHost:   ctx.String("mysqlhost"),
		DbUser:   ctx.String("user"),
		Database: ctx.String("database"),
	})

}
