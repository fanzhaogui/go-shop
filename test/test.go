package test

import (
	"shop/src/content/category"
	"fmt"
)

func Test_GetContentByIDDao()  {
	r := category.GetContentByIDDao(1)
	fmt.Println(r)
}