package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //导入驱动 ，MYSQL
	"fmt"
)

var (
	db *sql.DB
)

/**
  建立数据库连接
 */
func connect() {
	var er error
	db, er = sql.Open("mysql", "root:123456@tcp(localhost:3306)/world?charset=utf8")

	if er != nil {
		panic("db connect error")
	}

	if e := db.Ping(); e != nil {
		panic("ping error")
	}

}

/**
  单行查询
 */
func queryOne() {
	var n string
	var a int
	err := db.QueryRow("SELECT * FROM people WHERE name=?", "zhangcong").Scan(&n, &a)

	if err != nil {
		panic("queryone error")
	}

	fmt.Println(n, a)

}

/**
  查询多行记录
 */
func queryMany() {
	row, e := db.Query("SELECT * FROM people")
	if e != nil {
		panic("e")
	}
	defer row.Close()
	/**
	Scan()函数要求传递给它的目标变量的数目，与结果集中的列数正好匹配，否则就会出错。
    但总有一些情况，用户事先并不知道返回的结果到底有多少列，例如调用一个返回表的存储过程时。
    在这种情况下，使用rows.Columns()来获取列名列表。在不知道列类型情况下，
	应当使用sql.RawBytes作为接受变量的类型。获取结果后自行解析。
	 */
	rows,_:=row.Columns()
	for n,index:=range rows {
		fmt.Println(n,index)
	}
	for row.Next() {
		var n string
		var a int
		row.Scan(&n, &a)
		fmt.Println(n, a)
	}
}

/**
  添加记录
 */
func insert() {
	//方案一
	//r,e:=db.Exec("INSERT INTO people(name,age)VALUE (?,?)","wangcong",32)
	//if e!=nil {
	//	panic("insert error")
	//}
	//index,_:=r.RowsAffected()
	//fmt.Println(index)
	//方案二
	smtp, er := db.Prepare("INSERT INTO people(name,age)VALUE (?,?)")
	if er != nil {
		panic("insert error")
	}
	defer smtp.Close()
	res, e := smtp.Exec("zhanghaiquan", 44)
	if e != nil {
		panic(e)
	}
	fmt.Println(res.RowsAffected())
}

/**
  删除记录
 */
func delete() {

	smtp, err := db.Prepare("DELETE FROM people WHERE name=?")

	if err != nil {
		panic("delete error")
	}
	defer smtp.Close()

	res, _ := smtp.Exec("wangcong")
	fmt.Println(res.LastInsertId())

}

/*
  更新数据记录
 */

func update() {
	smtp, err := db.Prepare("UPDATE people set age=100 WHERE name=?")

	if err != nil {
		panic("update error")
	}
	defer smtp.Close()

	res, _ := smtp.Exec("zhanghaiquan")
	fmt.Println(res.RowsAffected())
}

/**
在事务中使用defer stmt.Close()是相当危险的。
因为当事务结束后，它会释放自己持有的数据库连接，
但事务创建的未关闭Stmt仍然保留着对事务连接的引用。
在事务结束后执行stmt.Close()，如果原来释放的连接已经被其他查询获取并使用，
就会产生竞争，极有可能破坏连接的状态。
 */
func commit() {
	tx, err := db.Begin()
	if err != nil {
		panic("commit error")
	}
	smtp, e := tx.Prepare("INSERT INTO people(name,age)VALUE (?,?)")
	if e != nil {
		panic("commit error smtp")
	}
	res, _ := smtp.Exec("zhaoguo", 21)
	fmt.Println(res.LastInsertId())
	//回滚
	defer tx.Rollback()
	//提交事务
	//tx.Commit()
}

func main() {
	connect()
	defer db.Close()

	queryMany()

	//queryOne()

	//insert()

	//delete()

	//update()

	//commit()
}
