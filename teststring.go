package main

import (
	"fmt"
	//"strings"
	//"strconv"
	"strings"
	//"reflect"
)

func main()  {
	//类型转换
	//字符串转int
	//var trans01 string = "300"
	//i,_ := strconv.Atoi(trans01)
	//fmt.Println(i)

	//字符串转为float
	//var trans02 string  = "290.22"
	//j,_ := strconv.ParseFloat(trans02,10)
	//fmt.Println(j)

	//int  转化为字符串
	//var tran03 int = 100
	//c := strconv.Itoa(tran03)
	//fmt.Println(c)


	//var str01 string = "nihao zhongguo"

	//字符串包含
	//fmt.Println(strings.Contains(str01,"nihao"))

	//字符串位置
	//fmt.Println(strings.Index(str01,"zh"))

	//最后出现的位置
	//fmt.Println(strings.LastIndex(str01,"o"))

	//统计字符串出现的次数
	//fmt.Println(strings.Count(str01,"o"))

	//替换
	//r := strings.Replace(str01,"o","**",1)
	//fmt.Println(r)

	//删除头部和尾部的字符串
	//fmt.Println(strings.TrimRight(r,"**"))

	//转化为大写
	//fmt.Println(strings.ToUpper(str01))

	//判断是否有前缀
	//fmt.Println(strings.HasPrefix(str01,"ni"))
	//fmt.Println(strings.HasSuffix(str01,"guo"))

	//根据空白分隔
	//fieldsStr := "  hello   it's  a  nice day today    "
	//
	//fieldsSpace  := strings.Fields(fieldsStr)
	//
	//fmt.Println(fieldsSpace)

	//根据指定字符串分隔
	slice01 := strings.Split("q,w,e,r,t,y,", ",")
	fmt.Println(slice01)

	//根据指定字符串拼接
	fmt.Println(strings.Join(slice01,"|"))











}
