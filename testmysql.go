package main

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"log"
	"fmt"
	"time"
)

func main()  {


	db,err := sql.Open("mysql","buzhisys:buzhisys@tcp(118.190.71.77:3306)/bz_admin?charset=utf8")
	if err != nil{
		log.Fatal("链接数据库错误")
	}
	////方法1:
	//result,err := db.Exec("INSERT INTO `bz_test` (`username`, `mobile`, `created_time`) VALUES (?, ?, ?)","zouying01","18932423859",time.Now().Unix())
	//result.LastInsertId()

	////方法2:
	//stmt,err :=  db.Prepare("INSERT `bz_test` SET `username`=?,`mobile`=?,`created_time`=?")
	//result , err := stmt.Exec("邹英02" , "18932423856" , time.Now().Unix())
	//
	//if err != nil{
	//	fmt.Println(err)
	//}
	//
	//id,err := result.LastInsertId()
	//if err != nil{
	//	fmt.Println(err)
	//}
	//fmt.Printf("插入的id为:%v",id)

	//
	//var username, mobile string
	//var created_time int


	rows,err := db.Query("SELECT `username`,`mobile`,`created_time` FROM `bz_test` WHERE `mobile`=?", "18932423859")
	if err != nil{
		log.Fatal("err")
	}
	for rows.Next(){
		var username, mobile string
		var created_time int
		if  err := rows.Scan(&username,&mobile,&created_time);err != nil{
			log.Fatal("数据没有查到")
		}
		fmt.Println(mobile)
		fmt.Println(username)
		fmt.Println(created_time)
	}












	
}
