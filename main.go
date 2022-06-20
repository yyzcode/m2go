package main

import (
	"github.com/urfave/cli/v2"
	"github.com/yyzcoder/m2go/app"
	"log"
	"os"
)

func main() {

	cliApp := &cli.App{
		Name:   "m2go",
		Usage:  "make structs from database information",
		Action: app.Run,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "overwrite",
				Aliases: []string{"w"},
				Value:   false,
				Usage:   "overwrite exist files",
			},
			&cli.StringFlag{
				Name:    "prefix",
				Aliases: []string{"p"},
				Value:   "",
				Usage:   "prefix of table name will be ignore in struct name and go file name",
			},
			&cli.StringFlag{
				Name:    "user",
				Aliases: []string{"u"},
				Value:   "root:123456",
				Usage:   "mysql username and password. example root:123456. if haven't pwd use root: ",
			},
			&cli.StringFlag{
				Name:    "mysqlhost",
				Aliases: []string{"host"},
				Value:   "127.0.0.1:3306",
				Usage:   "mysql host and port. example 127.0.0.1:3306",
			},
			&cli.StringFlag{
				Name:    "database",
				Aliases: []string{"db"},
				Usage:   "database name",
			},
			&cli.BoolFlag{
				Name:    "json",
				Aliases: []string{"j"},
				Value:   false,
				Usage:   "whether generate structs' json tag",
			},
			&cli.BoolFlag{
				Name:    "default_value",
				Aliases: []string{"dv"},
				Value:   false,
				Usage:   "whether generate field default value note",
			},
		},
	}
	if len(os.Args) == 1 {
		os.Args = append(os.Args, "--help")
	}
	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
