package other

import "encoding/json"

type Data struct {
	ID  *string         `json:"id"`
	Doc json.RawMessage `json:"doc"`
}

// ESIndexResponse ES数据 导出导入数据都可存储在里面
type ESIndexResponse struct {
	Data []Data `json:"data"`
}
