package main

import "fmt"

func main()  {

	//s := []int{1,2,5,6,9}
	//s := make([]int,10)
	//s1 := s[:]
	//
	//fmt.Println(append(s1,1))

	//map操作

	//var map1  map[string]string
	//map1  = make(map[string]string)
	//map1["cc"]  = "beijing"
	//map1["sss"] = "changsha"
	//
	////delete(map1,"cc")
	//v,ok  := map1["cc"]
	//
	//if ok {
	//	fmt.Println(v)
	//} else {
	//	fmt.Println("不存在")
	//}

	//for k := range map1{
	//	fmt.Println(map1[k])
	//}


	//for i, c := range "go,你好" {
	//	fmt.Printf("%d %c\n",i,c)
	//}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}


}
