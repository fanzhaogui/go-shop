package desc

import (
	"shop/src/commons"
	"fmt"
)

func InsertItemDescDao(t TbItemDesc) int {
	count, err := commons.Dml("insert into tb_item_desc values(?, ?, ?, ?)", t.ItemID, t.ItemDesc, t.Created, t.Updated)
	if err != nil {
		fmt.Println("insert item desc err: ", err)
		return -1
	}
	return int(count)
}