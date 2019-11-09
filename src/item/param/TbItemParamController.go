package param

import (
	"net/http"
	"shop/src/commons"
	"encoding/json"
	"strconv"
	"fmt"
)

// Router
func TbItemParamHandler() {
	// commons.Router.HandleFunc("/item/param/show", showParamController)
	// commons.Router.HandleFunc("/item/param/delete", delByIdsController)
	commons.Router.HandleFunc("/item/param/isCat", insertParamController)
	commons.Router.HandleFunc("/item/param/add", isCatController)
}

// 新增一条规格
func insertParamController(w http.ResponseWriter, r *http.Request)  {
	// 接收参数
	cid, _ := strconv.Atoi(r.FormValue("catid"))
	paramData := r.FormValue("param_data")

	s := InsertParamService(cid, paramData)
	// 返回参数
	b, _ := json.Marshal(s)
	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	w.Write(b)
}

// 分类对应的规格是否存在
func isCatController(w http.ResponseWriter, r *http.Request)  {
	catId, _ := strconv.Atoi(r.FormValue("cat_id"))
	er := selByCatIdDao(catId)
	b, _ := json.Marshal(er)
	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	w.Write(b)
}

// 删除规格参数 通过ids字符串
func delByIdsController(r *http.Request, w http.ResponseWriter)  {
	er := DelByIdsService(r.FormValue("ids"))
	b, _ := json.Marshal(er)
	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	w.Write(b)
}

// 分页列表
func showParamController(r *http.Request, w http.ResponseWriter)  {
	page, _ := strconv.Atoi(r.FormValue("page"))
	pageSize, _ := strconv.Atoi(r.FormValue("page_size"))
	fmt.Println(pageSize,page)
}