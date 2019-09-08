package commons

// 分页数据
type Datagrid struct {
	// 当前页显示的数据
	Rows interface{} `json:"rows"`
	// 总个数
	Total int `json:"total"`
}

// easy tree的字段
type EasyUITree struct {
	Id     int    `json:"id"`
	Text   string `json:"text"`
	State string `json:"state"`
}
