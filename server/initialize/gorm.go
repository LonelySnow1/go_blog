package initialize

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"server/global"
)

func InitGorm() *gorm.DB {
	mysqlCfg := global.Config.Mysql

	db, err := gorm.Open(mysql.Open(mysqlCfg.Dsn()), &gorm.Config{ //创建gorm链接对象
		Logger: logger.Default.LogMode(mysqlCfg.LogLevel()),
	})
	if err != nil {
		global.Log.Error("Failed to connect to mysql", zap.Error(err))
		os.Exit(1)
	}

	sqlDB, _ := db.DB()                          // 获取标准库对象
	sqlDB.SetMaxIdleConns(mysqlCfg.MaxIdleConns) // 最大空闲链接数
	sqlDB.SetMaxOpenConns(mysqlCfg.MaxOpenConns) // 最大打开链接数

	return db
}
