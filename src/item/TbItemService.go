package item

import (
	"shop/src/commons"
	"shop/src/item/cat"
	"strings"
	"mime/multipart"
	"io/ioutil"
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

// 上架
func instock(ids string) (e commons.EgoResult) {
	count := updateStatusByIdsDao(strings.Split(ids, ","), 1)
	if count > 0 {
		e.Status = 200
	}
	return
}

// 下架
func uninstock(ids string) (e commons.EgoResult) {
	count := updateStatusByIdsDao(strings.Split(ids, ","), 2)
	if count > 0 {
		e.Status = 200
	}
	return
}


// 保存文件
func imageUpdateService(f multipart.File, h *multipart.FileHeader) map[string]interface{}{
	m := make(map[string]interface{})
	b, err := ioutil.ReadAll(f)
	if err != nil {
		m["error"] = 1
		m["message"] = "读取上传文件信息失败, err: " + err.Error()
		return m
	}

	// 纳秒时间戳 + 随机数+ 扩展名
	// rand.Seed(time.Now().UnixNano())
	// fileName := strconv.Itoa(int(time.Now().UnixNano())) + strconv.Itoa(rand.Intn(1000))+ h.Filename[strings.
	// 	LastIndex(h.Filename,
	// 	"."):]
	fileName := "static/images/" + commons.GenerateFileName(h.Filename[strings.LastIndex(h.Filename,"."):])
	err = ioutil.WriteFile(fileName, b, 0777)
	if err != nil {
		m["error"] = 1
		m["message"] = "保存文件时失败，err: " + err.Error()
		return m
	}

	m["error"] = 0
	m["url"] = commons.CurrentPath + fileName
	return m

}
