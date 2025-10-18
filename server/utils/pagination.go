package utils

import (
	"server/global"
	"server/model/other"
)

// MySQLPagination 实现 MySQL 数据分页查询
func MySQLPagination[T any](model *T, option other.MySQLOption) (list []T, total int64, err error) {
	// 设置分页的默认值
	if option.Page < 1 {
		option.Page = 1 // 页码不能小于1，默认为1
	}
	if option.PageSize < 1 {
		option.PageSize = 10 // 每页记录数不能小于1，默认为10
	}
	if option.Order == "" {
		option.Order = "id desc" // 默认按id降序排列
	}

	// 创建查询
	query := global.DB.Model(model)

	// 如果传入了额外的 WHERE 条件，则应用这些条件
	if option.Where != nil {
		query = query.Where(option.Where)
	}

	// 计算符合条件的记录总数
	if err = query.Count(&total).Error; err != nil {
		return nil, 0, err // 如果查询总数失败，返回错误
	}

	// 预加载关联模型
	for _, preload := range option.Preload {
		query = query.Preload(preload) // 应用预加载的关联查询
	}

	// 应用分页查询
	err = query.Order(option.Order).
		Limit(option.PageSize).                      // 设置每页记录数
		Offset((option.Page - 1) * option.PageSize). // 设置偏移量，根据页码计算
		Find(&list).Error                            // 执行查询，并将结果存入 list 中

	return list, total, err // 返回分页结果和总记录数
}
