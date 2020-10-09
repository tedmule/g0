package main

import (
	"fmt"
	"regexp"
)

func formatServiceString(s string) string {
	/*
		Input string and return:
		abc         -> abc
		abc.com     -> abc-com
		abc..com    -> abc-com
		abc_com     -> abc-com
		c.com_cn  	-> abc-com-cn
	*/
	replaced := regexp.MustCompile(`\.+|_+`)
	return replaced.ReplaceAllString(s, "-")
}

func main() {
	fmt.Println(formatServiceString("abc.com"))
	fmt.Println(formatServiceString("abc.com"))
	fmt.Println(formatServiceString("abc...com"))
	fmt.Println(formatServiceString("abc_com"))
	fmt.Println(formatServiceString("abc_com.cn"))
}
