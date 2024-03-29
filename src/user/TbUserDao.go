package user

import (
	"shop/src/commons"
	"fmt"
	"log"
)

// 根据用户名和密码查询
//
func SelByUnPwdDao(un, pwd string) *TbUser {
	sql := "select id,username," +
		"password, phone, email, " +
		"created, updated from tb_user where username =? and password=? or email=? and password=? or phone" +
		"=? and password=?"
	rows, err := commons.Dql(sql, un, pwd, un, pwd, un, pwd)
	defer rows.Close()
	fmt.Println(sql, un, pwd)
	if err != nil {
		fmt.Println("err: ", err)
		return nil
	}

	if rows.Next() {
		user := new(TbUser)
		rows.Scan(&user)
		log.Println(user.Username)
		return user
	}
	return nil
}

func SelUserByAny(any string) TbUser {
	rows, err := commons.Dql("select * from tb_user where username=? or phone = ? or email = ?", any, any, any)
	defer rows.Close()
	var user TbUser
	if err != nil {
		log.Fatal(err)
		return user
	}

	if rows.Next() {
		rows.Scan(user)
	}
	fmt.Println("查询结果为空！")
	return user

}