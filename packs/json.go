package packs

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	Name   string
	Gender string
}

func DemoJson() {
	s := "hello"
	fmt.Println(reflect.TypeOf(s))
	fmt.Printf("%T\n", s)

	user := &User{Name: "Frank", Gender: "male"}
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

}
