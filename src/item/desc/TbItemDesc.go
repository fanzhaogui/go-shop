package desc

type TbItemDesc struct {
	ItemID   int    `json:"item_id"`
	ItemDesc string `json:"item_desc"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}