package service

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"server/global"
)

type EsService struct{}

// IndexCreate 创建索引
func (esService *EsService) IndexCreate(indexName string, mapping *types.TypeMapping) error {
	_, err := global.ESClient.Indices.Create(indexName).Mappings(mapping).Do(context.TODO())
	return err
}

// IndexDelete 删除索引
func (esService *EsService) IndexDelete(indexName string) error {
	_, err := global.ESClient.Indices.Delete(indexName).Do(context.TODO())
	return err
}

// IndexExist 判断索引是否存在
func (esService *EsService) IndexExist(indexName string) (bool, error) {
	return global.ESClient.Indices.Exists(indexName).Do(context.TODO())
}
