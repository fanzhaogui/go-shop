package cat

func ShowCatByIdService(id int) *TbItemCat {
	return selByIdDao(id)
}
