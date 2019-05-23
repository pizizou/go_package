package main

import (
	"fmt"
	"time"
)

func main() {
	var t time.Time // 定义 time.Time 类型变量
	t = time.Now()  // 获取当前时间
	//fmt.Printf("时间: %v, 时区:  %v,  时间类型: %T\n", t, t.Location(), t)
	//
	//// time.UTC() time 返回UTC 时区的时间
	fmt.Printf("时间: %v, 时区:  %v,  时间类型: %T\n", t.UTC(), t.UTC().Location(), t)
	//
	//now := time.Now()
	//m, _ := time.ParseDuration("-1m")
	//m1 := now.Add(m)
	//fmt.Println(m1)
	//
	//
	//h, _ := time.ParseDuration("-1h")
	//h1 := now.Add(8 * h)
	//fmt.Println(h1)
	//
	//subM := now.Sub(m1)
	//fmt.Println(subM.Minutes(), "分钟")
	//
	//
	//sumH := now.Sub(h1)
	//fmt.Println(sumH.String(), "小时")

	//m1,_ := time.ParseDuration("-1m")

	//fmt.Println(time.Now().Format("2006/01/02 15|04|05"))
	//
	//
	//time.Now()
	//
	//format := "2006-01-02 15:04:05"
	//a, _ := time.Parse(format, "2015-03-10 11:00:00")
	//b, _ := time.Parse(format, "2015-03-10 16:00:00")
	//fmt.Println(a.Unix())
	//fmt.Println(b.Unix())
	//fmt.Println(time.Unix(1425985200,0))

}