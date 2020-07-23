package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name   string
	Gender string
}

func main() {
	user := &User{Name: "Frank", Gender: "male"}
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
