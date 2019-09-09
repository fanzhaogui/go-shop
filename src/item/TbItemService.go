package item

import (
	"shop/src/commons"
	"shop/src/item/cat"
	"strings"
	"mime/multipart"
	"io/ioutil"
	"net/url"
	"strconv"
	"time"
	"shop/src/item/desc"
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

func insertTbItemService(f url.Values) (re commons.EgoResult) {
	var item TbItem
	id := commons.GenerateId()
	item.Id = id
	cid, _ := strconv.Atoi(f["Cid"][0])
	price, _ := strconv.Atoi(f["Price"][0])
	num, _ := strconv.Atoi(f["Num"][0])
	date := time.Now().Format("2006-01-02 15:04:05")
	item.Cid = cid
	item.Title = f["Title"][0]
	item.SellPoint = f["SellPoint"][0]
	item.Price = price
	item.Num = num
	item.Image = f["Image"][0]
	item.Updated = date
	item.Created = date
	item.Status = 1

	count := insertItemDao(item)

	if count > 0 {
		// 添加商品描述
		var descData desc.TbItemDesc
		descData.ItemID = id
		descData.ItemDesc = f["Desc"][0]
		descData.Created = date
		descData.Updated = date
		count = desc.InsertItemDescDao(descData)
		if count > 0 {
			re.Status = 200
		} else {
			// 删除添加商品信息
			delItemDao(id)
			re.Status = 400
		}
	} else {
		re.Status = 400
	}
	return re
}