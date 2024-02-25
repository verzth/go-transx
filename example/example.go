package main

import (
	"fmt"
	"git.verzth.work/go/transx"
)

func main() {
	a := A{Name: "John", Age: 20}
	b := B{}
	err := transx.Transform(a, &b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", b)
}

type A struct {
	Name string `transx:"name"`
	Age  int    `transx:"age"`
}

type B struct {
	Nama string `transx:"name"`
	Umur int    `transx:"age"`
}
