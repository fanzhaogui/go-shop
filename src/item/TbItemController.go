package item

import (
	"shop/src/commons"
	"net/http"
	"strconv"
	"encoding/json"
)

// 访问控制
func ItemHandler()  {
	// query
	commons.Router.HandleFunc("/showItem", showItemController)
	// delete
	commons.Router.HandleFunc("/item/delete", delByIdsController)
	// 上架，下架

}

// 显示商品信息
func showItemController(w http.ResponseWriter, r *http.Request)  {
	page, _ := strconv.Atoi(r.FormValue("page"))
	rows, _ := strconv.Atoi(r.FormValue("rows"))

	datagrid := showItemService(page, rows)
	b, _ := json.Marshal(datagrid)

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}

// 商品删除
func delByIdsController(w http.ResponseWriter, r *http.Request) {
	ids := r.FormValue("ids")

	er := delByIdsService(ids)

	b, _ := json.Marshal(er)

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}