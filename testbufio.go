package main

import (
	"strings"
	"fmt"
	"bufio"

)

func main() {
	r := strings.NewReader("hello world !")
	reader := bufio.NewReader(r)

	bytes, _ := reader.Peek(5)
	fmt.Printf("%s\n",bytes)
	//n, _ := reader.Read(bytes)
	//fmt.Println(n)
	//reader.Discard(1)


}