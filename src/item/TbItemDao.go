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