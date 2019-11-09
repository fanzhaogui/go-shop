package user

import (
	"shop/src/commons"
	"github.com/lunny/log"
)

func LoginService(un, pwd string) (re commons.EgoResult)  {
	u := SelByUnPwdDao(un, pwd)
	log.Info(u)
	if u != nil {
		re.Status = 200
	} else {
		re.Status = 400
	}
	return
}
