package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"io"
)

/**
主要测试文件的写入 和读取
文件操作可以参考网址：https://colobu.com/2016/10/12/go-file-operations/
 */

 //写入文件
func write()  {
	filename := "test.txt"
	fout,err := os.Create(filename)
	defer fout.Close()
	if err != nil{
		fmt.Println(err)
		return
	}
	for i:=0;i<10;i++{
		fout.WriteString("string:Just a test \r\n")
		fout.Write([]byte("bype:just a test\r\n"))
	}
}

//从文件中读取文件
func read()  {
	filename := "test.txt"
	fl , err := os.Open(filename)
	defer fl.Close()
	if err != nil{
		fmt.Println(err)
		return
	}
	buf := make([]byte,1024)
	for {
		n,_ := fl.Read(buf)
		if 0 == n{
			break
		}
		os.Stdout.Write(buf[:n])
	}
}

func readLine()  {
	filename := "test.txt"
	fl,err := os.OpenFile(filename,os.O_RDWR,0666)
	if err != nil{
		fmt.Println("Open file err:",err)
	}
	defer fl.Close()
	stat,err := fl.Stat()
	if err != nil{
		panic(err)
	}

	var size = stat.Size()
	fmt.Println("file size=",size)

	buf := bufio.NewReader(fl)

	for{
		line , err := buf.ReadString('\n')
		line = strings.TrimSpace(line)

		fmt.Println(line)
		if err != nil{
			if err == io.EOF{
				fmt.Println("FILE read ok!")
				break
			}else {
				fmt.Println("file read err:",err)
				return
			}
		}
	}

	
}

func main() {
	// 获取当前目录
	//dir, _ := os.Getwd()
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


	//file := dir + "new"
	//fi, _ := os.Stat(file)
	//if fi == nil {
	//	fh, _ := os.Create(file) // 文件不存在就创建
	//	fh.Write([]byte("package main\nfunc main(){\n}\n"))
	//} else {
	//	fh, _ := os.OpenFile(file, os.O_RDWR, 0666) // 文件存在就打开
	//	fh.Write([]byte("package main\nfunc main(){\n}\n"))
	//}

	//write()
	//read()

	readLine()





}