package main

import (
	"os"
	"fmt"
	"flag"
	"bufio"
)

/**
命令行解析测试
 */

func ParseOsArgs()  {
	if len(os.Args) < 2{
		fmt.Println("no args")
		return
	}
	fmt.Println("script name:",os.Args[0])

	for i := range os.Args{
		fmt.Printf("this is %d arg : %s \n",i,os.Args[i])
	}
}

func ParseFlag()  {
	//ip := flag.String("ip","127.0.0.1","this is you ip")
	var sip string
	var h bool
	flag.StringVar(&sip,"ip","127.0.0.1","this is you ip")
	flag.BoolVar(&h,"h",false,"mifdfd")

	flag.Parse()
	//fmt.Println(sip)
	//fmt.Println(h)

	if h {
		flag.PrintDefaults()
	}

	for i,param := range flag.Args(){
		fmt.Printf("#%d    :%s\n",i,param)
	}

}

func ParseStdin()  {

	//方式一
	var str1 string

	fmt.Print("Please input str1: ")
	fmt.Scanf("%s\n", &str1)
	fmt.Println("Output str1: ", str1)

	fmt.Println("----------------------------")

	var str2, str3 string
	fmt.Print("Please input str2-str3: ")
	fmt.Scanf("%s %s\n", &str2, &str3)

	fmt.Println("Output str2-str3: ", str2, str3)

	fmt.Println("----------------------------")

	var str4, str5 string
	fmt.Print("Please input str4-str5: ")
	fmt.Scanf("%s %s", &str4, &str5)

	fmt.Println("Output str4-str5: ", str4, str5)


	//方式二
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Println("用户名:")
	//username ,_ := reader.ReadString('\n')
	//fmt.Println("密码:")
	//password ,_ := reader.ReadString('\n')
	//
	//fmt.Println("用户名和密码:",username,password)


}

func main()  {
	//ParseOsArgs()
	//ParseFlag()
	//ParseStdin()



}

