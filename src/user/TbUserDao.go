package user

import (
	"shop/src/commons"
	"fmt"
)

// 根据用户名和密码查询
//
func SelByUnPwdDao(un, pwd string) *TbUser {
	sql := "select id,username," +
		"password, phone, email, " +
		"created, updated from tb_user where username =? and password=? or email=? and password=? or phone" +
		"=? and password=?"
	rows, err := commons.Dql(sql, un, pwd, un, pwd, un, pwd)
	if err != nil {
		fmt.Println("err: ", err)
		return nil
	}

	if rows.Next() {
		user := new(TbUser)
		rows.Scan(&user)
		return user
	}
	return nil
}