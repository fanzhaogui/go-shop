package item

import (
	"shop/src/commons"
	"net/http"
	"strconv"
	"encoding/json"
	"fmt"
)

// 访问控制
func ItemHandler()  {
	// query
	commons.Router.HandleFunc("/showItem", showItemController)
	// delete
	commons.Router.HandleFunc("/item/delete", delByIdsController)
	// 上架，下架
	commons.Router.HandleFunc("/item/instock", instockByIdsController)
	commons.Router.HandleFunc("/item/offstock", uninstockByIdsController)

	// 图片上传
	commons.Router.HandleFunc("/item/imageupload", imageUploadController)
}


// 图片上传 - 前端在上传多文件时，其实是多次文件上传
func imageUploadController(w http.ResponseWriter, r *http.Request)  {
	multipartFile, multipartFileHeader, err := r.FormFile("imgFile")
	fmt.Println("上传文件时出错", err)
	m := make(map[string]interface{})
	if err != nil {
		m["error"] = 1
		m["message"] = "接收图片错误，err : " + err.Error()
	} else {
		m = imageUpdateService(multipartFile, multipartFileHeader)
	}

	b, _ := json.Marshal(m)
	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	w.Write(b)
}

// 显示商品信息
func showItemController(w http.ResponseWriter, r *http.Request)  {
	page, _ := strconv.Atoi(r.FormValue("page"))
	rows, _ := strconv.Atoi(r.FormValue("rows"))

	datagrid := showItemService(page, rows)
	b, _ := json.Marshal(datagrid)

	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	w.Write(b)
}

// 商品删除
func delByIdsController(w http.ResponseWriter, r *http.Request) {
	ids := r.FormValue("ids")

	er := delByIdsService(ids)

	b, _ := json.Marshal(er)

	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	w.Write(b)
}

// 商品上架
func instockByIdsController(w http.ResponseWriter, r *http.Request) {
	ids := r.FormValue("ids")

	er := instock(ids)

	b, _ := json.Marshal(er)

	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	w.Write(b)
}

// 商品下架
func uninstockByIdsController(w http.ResponseWriter, r *http.Request) {
	ids := r.FormValue("ids")

	er := uninstock(ids)

	b, _ := json.Marshal(er)

	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	w.Write(b)
}