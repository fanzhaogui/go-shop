package cat

type TbItemCat struct {
	Id        int    `json:"id"`
	ParentId  int    `json:"parent_id"`
	Name      string `json:"name"`
	Status    byte   `json:"status"`
	SortOrder int    `json:"sort_order"`
	IsParent  byte   `json:"is_parent"`
	Created   string `json:"created"`
	Updated   string `json:"updated"`
}
