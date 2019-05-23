package main

import (
	"os"

)

func main() {
	// 获取当前目录
	dir, _ := os.Getwd()
	//fmt.Println(dir, err)
	//file := dir + "/new"
	//var fh *os.File
	//fi, _ := os.Stat(file)
	//if fi == nil {
	//	fh, _ = os.Create(file) // 文件不存在就创建
	//} else {
	//	fh, _ = os.OpenFile(file, os.O_RDWR, 0666) // 文件存在就打开
	//}
	//w := []byte("hello go language" + time.Now().String())
	//n, err := fh.Write(w)
	//fmt.Println(n, err)
	//// 设置下次读写位置
	//ret, err := fh.Seek(0, 0)
	//fmt.Printf("%s %v %v\n","当前文件指针位置", ret, err)
	//b := make([]byte, 128)
	//n, err = fh.Read(b)
	//fmt.Printf("%d %v %s\n",n, err, string(b))
	//fh.Close()


	file := dir + "new"
	fi, _ := os.Stat(file)
	if fi == nil {
		fh, _ := os.Create(file) // 文件不存在就创建
		fh.Write([]byte("package main\nfunc main(){\n}\n"))
	} else {
		fh, _ := os.OpenFile(file, os.O_RDWR, 0666) // 文件存在就打开
		fh.Write([]byte("package main\nfunc main(){\n}\n"))
	}





}