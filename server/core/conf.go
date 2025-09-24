package core

import (
	"gopkg.in/yaml.v3"
	"log"
	"server/config"
	"server/utils"
)

// InitConf 从 YAML 文件加载配置
func InitConf() *config.Config {
	c := &config.Config{}
	yamlConf, err := utils.LoadYAML() // 读取yaml数据为字节数组
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err) //log.Fatalf相当于 log.Printf + os.Exit(1) 打印错误并终止
	}
	err = yaml.Unmarshal(yamlConf, c) // 反序列化，将数据写入c
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML configuration: %v", err)
	}
	return c
}
