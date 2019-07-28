package controller

import (
	"github.com/astaxie/beego"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type ShopController struct {
	beego.Controller
}

var Db *sqlx.DB

func init() {
	// 数据库的链接  用户名:密码@连接方式(地址:端口)/数据库名
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_db")
	if err != nil {
		fmt.Println("connect database err: ", err)
		return
	}

	Db = database
}

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Intro        string  `json:"intro"`
	Price        float64 `json:"price"`
	Price_yh     float64 `json:"price_yh"`
	Photo_little string  `json:"photo_little"`
	Saledcount   int     `json:"saledcount"`
}

// 商品列表的展示
//
func (this *ShopController) GetProductList() {
	cat_id, err := this.GetInt("cat_id")
	if err != nil {
		fmt.Println("get cate_id err :", err)
		return
	}

	var product []Product

	err = Db.Select(&product, "select id,name,intro,price,price_yh,photo_little, saledcount from m_product where cid= ?", cat_id)
	if err != nil {
		fmt.Println("select from product err: ", err)
		return
	}
	this.Data["json"] = product
	this.ServeJSON()
}

// 购物车
// 
type CartList struct {
	Id          int     `json:"id"`
	Num         int     `json:"num"`
	Uid         int     `json:"uid"`
	Pid         int     `json:"pid"`
	Price       float64 `json:"price"`
	Photo_little string  `json:"photo_little"` // 这里定义字段的方式注意
	Pro_name     string  `json:"pro_name"`
}

// 获取购物车中的数据
//
func (this *ShopController) GetCartList() {
	userId, err := this.GetInt("user_id")
	if err != nil {
		fmt.Println("get user_id err:", err)
		return
	}

	var cartList []CartList
	err = Db.Select(&cartList, "select c.id, num, uid, pid, c.price, p.photo_little, p.`name` pro_name from m_cart c left join m_product p ON c.pid=p.id where uid=?", userId)
	if err != nil {
		fmt.Println("get carlist err: ", err)
		return
	}

	this.Data["json"] = cartList
	this.ServeJSON()
}
