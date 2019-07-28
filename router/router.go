package router

import (
	"github.com/astaxie/beego"
	"shop/controller"
)

// setting http router
func init()  {
	beego.Router("/productlist", &controller.ShopController{}, "*:GetProductList")
	beego.Router("/cartlist", &controller.ShopController{}, "*:GetCartList")


}
