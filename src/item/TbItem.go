package item

// 商品信息
// type TbItem struct {
// 	Id        int    `json:"id"`
// 	Title     string `json:"title"`
// 	SellPoint string `json:"sell_point"`
// 	Price     int    `json:"price"`
// 	Num       int    `json:"num"`
// 	Barcode   string `json:"barcode"`
// 	Image     string `json:"image"`
// 	Cid       int    `json:"cid"`
// 	Status    int    `json:"status"`
// 	Created   string `json:"created"`
// 	Updated   string `json:"updated"`
// }

type TbItem struct {
	Id        int
	Title     string
	SellPoint string
	Price     int
	Num       int
	Barcode   string
	Image     string
	Cid       int
	Status    int
	Created   string
	Updated   string
}

// type TbItemChild struct {
// 	TbItem
// 	CategoryName string `json:"category_name"`
// }

type TbItemChild struct {
	TbItem
	CategoryName string
}