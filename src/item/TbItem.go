package item

// 商品信息
type TbItem struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	SellPoint string `json:"sell_point"`
	Price     int    `json:"price"`
	Num       int    `json:"num"`
	Barcode   string `json:"barcode"`
	Image     string `json:"image"`
	Cid       int    `json:"cid"`
	Status    int    `json:"status"`
	Created   string `json:"created"`
	Updated   string `json:"updated"`
}

// 给页面使用，实现商品类目
type TbItemChild struct {
	TbItem
	CategoryName string `json:"category_name"`
}

// 给修改页面使用
type TbItemDescChild struct {
	TbItem
	CategoryName string `json:"category_name"`
	Desc         string `json:"desc"`
}
