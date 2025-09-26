package flag

import (
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
)

func Run(c *cli.Context) {
	switch {
	case c.Bool("sql"):
		if err := SQL(); err != nil {
			global.Log.Error("Failed to initialize database structure:", zap.Error(err))
			return
		} else {
			global.Log.Info("Successful database structure initialized")
		}
	}
}

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "go_blog"
	app.Flags = []cli.Flag{
		sqlFlag,
	}
	app.Action = Run
	return app
}

func InitFlag() {
	app := NewApp()
	err := app.Run(os.Args)
	if err != nil {
		global.Log.Error("Failed to initialize database structure", zap.Error(err))
		os.Exit(1)
	}
}
