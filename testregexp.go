package main

import (
	"regexp"
	"fmt"
)

var mailRegexp = regexp.MustCompile(`([A-Za-z0-9]+)@([A-Za-z0-9.]+)\.([A-Za-z0-9.]+)`)

func main() {
	s := " hello's email is hellogo@gmail.com"
	matches := mailRegexp.FindStringSubmatch(s)

	fmt.Println(matches)// [hellogo@gmail.com hellogo gmail com]
	fmt.Println(matches[0]) // hellogo@gmail.com
	fmt.Println(matches[1]) // hellogo
	fmt.Println(matches[2]) // gmail
	fmt.Println(matches[3]) // com
}