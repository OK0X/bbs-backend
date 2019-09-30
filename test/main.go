package test

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	engine *xorm.Engine
)

func main() {
	//数据库连接参数
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "root", "root", "localhost:3306", "ok0xbbs")
	var err error
	//连接数据库
	engine, err = xorm.NewEngine("mysql", params)
	if err != nil {
		panic(err)
	}

	//执行sql查询，返回数据格式：[]map[string][]byte
	gsql := "SELECT `name` FROM `role`"
	gres, gerr := engine.Query(gsql)
	if gerr != nil {
		panic(gerr)
	}
	for _, v := range gres {
		fmt.Printf("%s\n", string(v["name"]))
	}

	// //执行sql命令(Insert/Update/Delete)
	// esql := "Update `go_member` SET `member_name` = ? WHERE `member_id` = ?"
	// eres, eerr := engine.Exec(esql, "tom", 4)
	// if eerr != nil {
	// 	panic(eerr)
	// }
	// num, ererr := eres.RowsAffected()
	// if ererr != nil {
	// 	panic(ererr)
	// }
	// fmt.Printf("总共修改了 %d 条数据\n", num)
}
