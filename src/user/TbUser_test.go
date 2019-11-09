package user

import (
	"testing"
	"fmt"
	"shop/src/commons"
)

func TestSelByUnPwdDao(t *testing.T) {
	rs := SelByUnPwdDao("tidy", commons.MD5("123456"))
	// log.Fatal("err")
	fmt.Println(rs.Username)
}

func TestLoginService(t *testing.T) {
	hexPws := commons.MD5("123456")
	re := LoginService("tidy", hexPws)
	fmt.Println("This ressult is : " , re.Data)
}

func TestSelUserByAny(t *testing.T) {
	s := SelUserByAny("tidy")

	t.Log(s)
}