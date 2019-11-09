package param

import (
	"shop/src/commons"
	"fmt"
)

/**
新增
 */
func insertParamDao(param TbItemParam) int {
	sql := "insert into tb_item_param values(default, ?, ?, ?, ?)"
	count, err := commons.Dml(sql, param.ItemCatId, param.ParamData, param.Created, param.Updated)

	if err != nil {
		fmt.Println("Insert TbItemParam err : ", err)
		return  -1;
	}
	return int(count)
}


/**
通过类目ID，查询对应的商品规格是否存在
 */
func selByCatIdDao(catid int) *TbItemParam {
	sql := "select * from tb_item_param where item_cat_id = ?"
	rows, err := commons.Dql(sql, catid)

	if err != nil {
		return nil
	}

	if rows.Next() {
		t := new (TbItemParam)
		// rows.Scan(&t.Id, &t.ItemCatId, &t.ItemCatId, &t.ParamData, &t.Created, t.Updated)
		rows.Scan(&t)
		return t
	}
	return  nil
}

// 通过ids字符串，删除
func delByIdsDao(ids string) int {
	count, err := commons.Dml("delete from tb_item_param where id in (?)", ids)
	if err != nil {
		return  -1;
	}
	return int(count)
}

/**
更新规格表
 */
func updateByIdDao(p TbItemParam) int {
	count, err := commons.Dml("update tp_item_param set item_cat_id = ?, param_data = ?, update_time=now()",
		p.ItemCatId, p.ParamData)
	if err != nil {
		return -1;
	}
	return int(count)
}

// 列表
func listDao(page, pageSize int) []TbItemParam {
	current := (page - 1) * pageSize
	r, err := commons.Dql("select * from tb_item_param limit ?,?", current, pageSize)
	ips := make([]TbItemParam, 0)
	if err != nil {
		return  ips
	}
	if r.Next() {
		var ip TbItemParam
		r.Scan(&ip)
		ips = append(ips, ip)
	}
	return ips
}

// 总条数
func totalDao() int {
	var total int
	r, err := commons.Dql("select count(*) from tb_item_param")
	if err == nil {
		if r.Next() {
			r.Scan(&total)
		}
	}
	return total
}