package main

import (
	"io/ioutil"
	"fmt"
	"strings"
)


//读取目录
func ReadDir()  {
	rd,_ := ioutil.ReadDir(".")
	for _,fi := range rd{
		if fi.IsDir() {
			//fmt.Printf("[%s]\n", fi.Name())
		} else {
			fmt.Println(fi.Name())
		}
	}
}

// 从文件中读取数据
func ReadAll()  {
	rd := strings.NewReader("12345678")
	data,_ := ioutil.ReadAll(rd)
	fmt.Println(string(data))
}

func ReadFile()  {
	data,_ := ioutil.ReadFile("test.txt")
	fmt.Println(string(data))
}

func WriteFile(data []byte)  {
	ioutil.WriteFile("wite.txt",data,0666)
}


func main()  {
//ReadDir()
//	ReadAll()
//	ReadFile()
	WriteFile([]byte("dreambee nihao"))
}