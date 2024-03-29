package item

import (
	"shop/src/commons"
	"fmt"
	"database/sql"
)

/**
rows：每页显示的条数
page: 当前第几页
 */
func selByPageDao(rows, page int) []TbItem {
	r, err := commons.Dql("select * from tb_item limit ?,?", rows*(page-1), rows)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	items := make([]TbItem, 0)
	// r.Scan(&items)
	for r.Next() {
		var t TbItem
		var s sql.NullString
		r.Scan(&t.Id, &t.Title, &t.SellPoint, &t.Price, &t.Num, &s, &t.Image, &t.Cid, &t.Status, &t.Created, &t.Updated)
		t.Barcode = s.String

		items = append(items, t)
	}
	fmt.Println(items)
	return items
}

// 总条数
func selCount() (count int)  {
	rows, err := commons.Dql("select count(*) from tb_item")

	if err != nil {
		fmt.Println(err)
		return -1
	}

	rows.Next()
	rows.Scan(&count)
	return
}

/**
返回值小于0，则表示失败
 */
func updateStatusByIdsDao(ids []string, status int) int {
	if len(ids) <= 0 {
		return -1
	}

	Dsql := "update tb_item set status = ? where "

	for i := 0; i < len(ids); i ++ {
		Dsql += " id = " + ids[i]
		if i < len(ids) - 1 {
			Dsql += " OR "
		}
	}

	count, err := commons.Dml(Dsql, status)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return int(count)
}


// 商品添加
func insertItemDao(t TbItem) int {
	count, err:= commons.Dml("insert into tb_item values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", t.Id, t.Title,t.SellPoint, t.Price, t.Num, t.Barcode, t.Image, t.Cid, t.Status, t.Created, t.Updated)

	if err != nil {
		fmt.Println("insert item err:", err)
		return -1
	}
	return int(count)

}

// 根据ID删除
func delItemDao(id int) int {
	count, err := commons.Dml("delete from tb_item where id = ?", id)
	if err != nil {
		fmt.Println("del item err: ", err)
	}
	return int(count)
}

// 描述
func selByIdDao(id int) (t *TbItem) {
	tr, err := commons.Dql("select * from tb_item where id = ?", id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// 关于值为NULL的处理
	if tr.Next() {
		tr.Scan(&t)
		return
	}
	return nil
}