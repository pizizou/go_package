package main

import (
	"os"
	"fmt"
	"flag"
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

func main()  {
	//ParseOsArgs()
	ParseFlag()
}
