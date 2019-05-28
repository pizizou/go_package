package main

import (
	"fmt"
	"math/rand"
	"time"
)

//获取随机字符串
func GetRandomString(num int) string  {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := 0
	for i< num{
		rand_num := r.Intn(len(bytes))
		result = append(result,bytes[rand_num])
		i++
	}
	return string(result)
}



func main()  {
	fmt.Println(GetRandomString(8))
}