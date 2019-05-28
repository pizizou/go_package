package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func md5v1(str string)  string{
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func md5v2(str string) string  {
	data := []byte(str)
	has := md5.Sum(data)
	md5str  := fmt.Sprintf("%x",has)
	return md5str
}

func md5v3(str string) string  {
	h := md5.New()
	io.WriteString(h,str)
	md5str  := hex.EncodeToString(h.Sum(nil))
	return  md5str
}

func main()  {
	str := "MD5testing"
	fmt.Println(md5v1(str))
	fmt.Println(md5v2(str))
	fmt.Println(md5v3(str))

}