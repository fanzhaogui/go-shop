package router

import (
	"github.com/astaxie/beego"
	"shop/controller"
)

// setting http router
func init()  {
	beego.Router("/productlist", &controller.ShopController{}, "*:GetProductList")
	beego.Router("/cartlist", &controller.ShopController{}, "*:GetCartList")
	beego.Router("/buy_cart", &controller.ShopController{}, "*:BuyCart")

	beego.Router("/getlist", &controller.ShopController{}, "*:GetCateProductList")


}
