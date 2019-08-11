package commons

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

var (
	db *sql.DB
	stmt *sql.Stmt
	rows *sql.Rows
)

// 打开数据库
//
func openConnec() (err error) {
	// 此处为等号，否则创建局部变量
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_db")

	if err != nil {
		fmt.Println("链接数据库时出错：", err)
		return
	}
	return
}


// 关闭数据库
//
func CloseConn()  {
	if rows != nil {
		db.Close()
	}

	if stmt != nil {
		stmt.Close()
	}

	if db != nil {
		db.Close()
	}
}

//  执行DML新增， 删除， 修改操作
//
func Dml(sql string, args ... interface{}) (int64, error) {
	err := openConnec()
	defer CloseConn()
	if err != nil {
		fmt.Println("执行DML时出错，打开连接失败， err :", err)
		return 0, err
	}

	stmt, err = db.Prepare(sql)
	if err != nil {
		fmt.Println("执行DML时出错，预处理，err:", err)
		return 0, err
	}

	result, err := stmt.Exec(args...)
	if err != nil {
		fmt.Println("执行DML时出错，执行SQL，err:", err)
		return 0, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		fmt.Println("执行DML时出错，获取受影响的行数，err:", err)
		return 0, err
	}
	return count, err
}

// 执行DQL查询
//
func Dql(sql string, args ...interface{}) (*sql.Rows, error) {
	err := openConnec()
	defer CloseConn()
	if err != nil {
		fmt.Println("执行DML时出错，打开连接失败， err :", err)
		return nil, err
	}

	stmt, err = db.Prepare(sql)
	if err != nil {
		fmt.Println("执行DML时出错，预处理，err:", err)
		return nil, err
	}
	// 此处参数是切片
	rows, err := stmt.Query(args...)
	if err != nil {
		fmt.Println("执行DML时出错，执行SQL，err:", err)
		return nil, err
	}
	return rows, nil
}