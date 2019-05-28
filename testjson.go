package main

import (
	"encoding/json"
	"fmt"
)
/**
主要测试 json转map  json转struct struct转json的操作
 */

type Server struct {
	ServerName string `json:"server_name,string"`
	ServerIp string	`json:"server_ip,	omitempty"`
}

type Serverslice struct {
	Servers []Server `json:"servers"`
}

//json字符串转化为struct  注意  结构体重定义的变量首字母必须为大写
func StrToStruct(data string)  {
	var s Serverslice

	err := json.Unmarshal([]byte(data),&s)

	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Println(s.Servers[0].ServerName)
}




//只能解析map
func StrToMap(data string)  {
	m := make(map[string]string)
	err := json.Unmarshal([]byte(data),&m)

	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(m)

	for k,v := range m{
		fmt.Println(k,":",v)
	}
}


func StructToStr()  {
	var s Serverslice
	s.Servers = append(s.Servers,Server{ServerName:"zouying1",ServerIp:"127.0.0.1"})
	s.Servers = append(s.Servers,Server{ServerName:"zouying2",ServerIp:"127.0.0.2"})
	b,err := json.Marshal(s)

	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(b))

}


func main()  {

	//解析数据到结构体重
//	str := `{"Servers":
//[{"servername":"buzhi","serverIp":"127.0.0.1"},
//{"serverName":"dreambee","serverIp":"127.0.0.2"}]}`

//StrToStruct(str)

//b := `{"IP": "127.0.0.1", "name": "SKY"}`
//StrToMap(b)

	StructToStr()





}



