package main

import (
	"fmt"
<<<<<<< HEAD
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
=======

	"github.com/g0gogo/packs"
)

func main() {
	fmt.Println("hello")
	packs.DemoJson()
>>>>>>> 50c081d3915ea6347322571587fa4504c87cd47b
}
