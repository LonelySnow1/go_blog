package flag

import (
	"fmt"
	"github.com/urfave/cli"
	"go.uber.org/zap"
	"os"
	"server/global"
)

var (
	sqlFlag = &cli.BoolFlag{
		Name:  "sql",
		Usage: "Initializes the structure of the database table",
	}
	sqlExportFlag = &cli.BoolFlag{
		Name:  "sql-export",
		Usage: "Exports SQL data to a specified file.",
	}
)

func Run(c *cli.Context) {
	// 拦截
	if c.NumFlags() > 1 {
		err := cli.NewExitError("Only one command can be specified", 1)
		global.Log.Error("Invaild command usage:", zap.Error(err))
		os.Exit(1)
	}
	switch {
	case c.Bool(sqlFlag.Name):
		if err := SQL(); err != nil {
			global.Log.Error("Failed to initialize database structure:", zap.Error(err))
			return
		} else {
			global.Log.Info("Successful database structure initialized")
		}
	case c.Bool(sqlExportFlag.Name):
		if err := SQLExport(); err != nil {
			global.Log.Error("Failed to export SQL data:", zap.Error(err))
		} else {
			global.Log.Info("Successfully exported SQL data")
		}
	default:
		err := cli.NewExitError("Unknow command", 1)
		global.Log.Error("Unknow command usage:", zap.Error(err))
	}
}

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "go_blog"
	app.Flags = []cli.Flag{
		sqlFlag,
		sqlExportFlag,
	}
	app.Action = Run
	return app
}

func InitFlag() {
	if len(os.Args) > 1 { //命令行的参数数量 也就是go run main.go -XXX 后的 -XXX数量
		app := NewApp()
		err := app.Run(os.Args)
		if err != nil {
			global.Log.Error("Failed to initialize database structure", zap.Error(err))
			os.Exit(1)
		}
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			fmt.Println("Display help message...")
		}
		os.Exit(0)
	}
}
