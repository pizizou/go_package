package main

import (
	//"log"
	//"log"
	//"os"
	"test/logger"
	"time"

	"log"
)

func main()  {

	//Print 接口
	//log.Print("fdfd","11")
	//log.Println("fadd","11")
	//log.Printf("my name is %v","zouying")

	//Fatal 接口
	//defer func() {
	//	fmt.Println("zouying")
	//}()
	//
	//log.Fatal("邹英")
	//log.Fatalf("my name is %v","zouying")
	//log.Fatalln("mydfd")

	//Panic 接口

	//level := "Panic"
	//defer func() {
	//	fmt.Println("defer Panic 1")
	//	if err := recover(); err != nil {
	//		fmt.Println(err)
	//	}
	//}()
	//log.Panic(level, " level")
	//defer func() {
	//	fmt.Println("defer Panic 2")
	//}()


	/**
	Ldate 显示当前日期（当前时区）
	Ltime 显示当前时间（当前时区）
	Lmicroseconds 显示当前时间（微秒）
	Llongfile 包含路径的完整文件名
	Lshortfile 不包含路径的文件名
	LUTC Ldata和Ltime使用UTC时间
	LstdFlags 标准Logger的标识，等价于 Ldate | Ltime
	 */


	//prefix := "[INFO]"
	//logger := log.New(os.Stdout, prefix, log.LstdFlags | log.Lshortfile)
	//logger.Print("Hello Go")
	//logger.SetPrefix("[OUTPUT]")
	//logger.SetFlags(log.LstdFlags)
	//logger.Print("Hello Logger")

	//logger.Config()


	logger.Infof("fdfddf")

	time.Sleep(5*time.Second)



}