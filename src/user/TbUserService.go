package user

import "shop/src/commons"

func LoginService(un, pwd string) (re commons.EgoResult)  {
	u := SelByUnPwdDao(un, pwd)

	if u != nil {
		re.Status = 200
	} else {
		re.Status = 400
	}
	return
}
