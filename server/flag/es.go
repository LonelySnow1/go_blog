package flag

import (
	"bufio"
	"fmt"
	"os"
	"server/elasticsearch"
	"server/service"
)

func Elasticsearch() error {
	esService := service.ServiceGroupApp.EsService

	indexExist, err := esService.IndexExists(elasticsearch.ArticleIndex())
	if err != nil {
		return err
	}

	if indexExist {
		fmt.Println("the index already exists. Do you want to delete and recreate the index?(y/n)")
		// 读取用户输入
		scanner := bufio.NewScanner(os.Stdin) // 创建一个扫描器，绑定到标准输入流，准备读取用户的键盘输入。
		scanner.Scan()                        // 执行一次读取操作（默认按「行」读取，直到遇到换行符 \n）
		input := scanner.Text()               // 这行代码将读取到的内容赋值给变量 input（不包含 \n）

		switch input {
		case "y":
			fmt.Println("Proceeding to delete the data and recreate the index")
			if err := esService.IndexDelete(elasticsearch.ArticleIndex()); err != nil {
				return err
			}
		case "n":
			fmt.Println("Exiting the program")
			os.Exit(0)
		default:
			fmt.Println("Invalid input.Please enter 'y' to delete and recreate the index,or 'n' to exit")
			return Elasticsearch()
		}
	}

	return esService.IndexCreate(elasticsearch.ArticleIndex(), elasticsearch.ArticleMapping()) // 索引和映射
}
