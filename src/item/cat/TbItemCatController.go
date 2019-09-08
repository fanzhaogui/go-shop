package cat

import (
	"net/http"
	"strconv"
	"encoding/json"
	"shop/src/commons"
)

// 访问控制
func CatHandler()  {
	// 新增展示的类目
	commons.Router.HandleFunc("/item/cat/show", ShowItemCatController)
}

/**
显示 类目
 */
func ShowItemCatController(w http.ResponseWriter, r *http.Request)  {
	id := r.FormValue("id")
	if id == "" {
		id = "0"
	}
	idInt, _ := strconv.Atoi(id)

	t := showParentCatSercie(idInt)

	b, _ := json.Marshal(t)
	w.Header().Set(commons.HEADER_CONTENT_TYPE, commons.JSON_HEADER)
	w.Write(b)
}
