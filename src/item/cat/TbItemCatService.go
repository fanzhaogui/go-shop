package cat

import "shop/src/commons"

func ShowCatByIdService(id int) *TbItemCat {
	return selByIdDao(id)
}

func showParentCatSercie(pid int) (tree []commons.EasyUITree) {
	pcats := selByPidDao(pid)

	tree = make([]commons.EasyUITree, 0)

	for _, v := range pcats {
		state := "open"
		if v.IsParent > 0 {
			state = "closed"
		}
		tree = append(tree, commons.EasyUITree{v.Id, v.Name, state})
	}
	return
}