package cat

import (
	"shop/src/commons"
	"fmt"
)

func selByIdDao(id int) (t *TbItemCat) {
	rows, err := commons.Dql("select * from tb_item_cat where id = ?", id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//
	if rows.Next() {
		t =new (TbItemCat)
		// rows.Scan(&t)
		rows.Scan(&t.Id,&t.ParentId,&t.Name,&t.Status,&t.SortOrder,&t.IsParent,&t.Created,&t.Updated)
	}
	return
}
