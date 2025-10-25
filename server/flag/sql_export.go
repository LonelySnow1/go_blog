package flag

import (
	"fmt"
	"os"
	"os/exec"
	"server/global"
	"time"
)

// SQLExport 导出 MySQL 数据
func SQLExport() error {
	mysqlCfg := global.Config.Mysql

	timer := time.Now().Format("20060102")
	sqlPath := fmt.Sprintf("mysql_%s.sql", timer)
	cmd := exec.Command("docker", "exec", "mysql", "mysqldump", "-u"+mysqlCfg.Username, "-p"+mysqlCfg.Password, mysqlCfg.DBName) // 本地运行
	//cmd := exec.Command("ssh", "ubuntu@203.0.113.5", "docker exec mysql mysqldump -u"+mysqlCfg.Username+" -p"+mysqlCfg.Password+" "+mysqlCfg.DBName)	 //服务器运行示例

	outFile, err := os.Create(sqlPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	cmd.Stdout = outFile
	return cmd.Run()
}
