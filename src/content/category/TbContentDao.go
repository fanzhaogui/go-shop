package category

import (
	"shop/src/commons"
	"fmt"
)

func GetContentByIDDao(id int) *TbContent {
	sql := "select * from tb_content where id = ?"
	rows, err := commons.Dql(sql, id)
	fmt.Println(sql, id)
	if err != nil {
		fmt.Println("err: ", err)
		return nil
	}

	if rows.Next() {
		c := new(TbContent)
		rows.Scan(&c)
		return c
	}
	return nil
}
