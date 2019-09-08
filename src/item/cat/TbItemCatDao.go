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

/**
父 分类
@param int pid
@return []TbItemCat
 */
func selByPidDao(pid int) (t []TbItemCat) {
	rows, err := commons.Dql("select * from tb_item_cat where parent_id = ?", pid);

	if err != nil {
		fmt.Println("selec parent cat err: ", err)
		return nil
	}

	var c = make([]TbItemCat, 0)
	for rows.Next() {
		var t TbItemCat
		rows.Scan(&t.Id,&t.ParentId,&t.Name,&t.Status,&t.SortOrder,&t.IsParent,&t.Created,&t.Updated)
		c = append(c, t)
	}
	return c
}