package item

import (
	"shop/src/commons"
	"shop/src/item/cat"
	"strings"
)

// 数据分页展示
// func showItemService(page, rows int) (e *commons.Datagrid) {
// 	ts := selByPageDao(rows, page)
// 	if ts != nil {
// 		e = new (commons.Datagrid)
//
// 		count := selCount()
// 		e.Total = count
//
// 		e.Rows = ts
// 		return
// 	}
//
// 	return nil
// }

func showItemService(page, rows int) (e *commons.Datagrid) {
	ts := selByPageDao(rows, page)
	if ts != nil {
		itemChildren := make([]TbItemChild, 0)
		for i := 0; i < len(ts); i++ {
			var itemChild TbItemChild
			itemChild.Id = ts[i].Id
			itemChild.Updated = ts[i].Updated
			itemChild.Created = ts[i].Created
			itemChild.Status = ts[i].Status
			itemChild.Barcode = ts[i].Barcode
			itemChild.Price = ts[i].Price
			itemChild.Num = ts[i].Num
			itemChild.SellPoint = ts[i].SellPoint
			itemChild.Title = ts[i].Title
			itemChild.CategoryName = cat.ShowCatByIdService(ts[i].Cid).Name

			itemChildren = append(itemChildren, itemChild)
		}

		e = new(commons.Datagrid)
		e.Rows = itemChildren
		e.Total = selCount()

		return
	}
	return nil
}

// 删除商品， 实际是更改商品状态为 status = 3
func delByIdsService(ids string) (e commons.EgoResult) {
	count := updateStatusByIdsDao(strings.Split(ids, ","), 3)
	if count > 0 {
		e.Status = 200
	}
	return
}