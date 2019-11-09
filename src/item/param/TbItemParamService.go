package param

import (
	"shop/src/commons"
	"time"
	"fmt"
)

/**
根据类目ID查询规格参数是否已经添加
 */
func CatidService(catid int) (e commons.EgoResult) {
	p := selByCatIdDao(catid)
	fmt.Println(p)
	if p == nil {
		e.Status = 200
	}
	return
}

// 新增规格参数模板
func InsertParamService(catid int, paramData string) (e commons.EgoResult) {
	timeLayout := commons.GetTimeLayout("default")
	fmt.Println("layout : " ,timeLayout)
	data := time.Now().Format(timeLayout)
	param := TbItemParam{ItemCatId: catid, ParamData: paramData, Created: data, Updated: data}
	count := insertParamDao(param)
	if count > 0 {
		e.Status = 200
	}
	return
}

// 删除多个
func DelByIdsService(ids string) (e commons.EgoResult) {
	count := delByIdsDao(ids)
	if count > 0 {
		e.Status = 200
	}
	return
}