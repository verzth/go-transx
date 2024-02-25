package main

import (
	"encoding/json"
	"fmt"
	"git.verzth.work/go/transx"
)

func main() {
	transform()
	transformSlice()
}

func transform() {
	a := A{Name: "John", Age: 20}
	b := B{}
	err := transx.Transform(a, &b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", b)
}

func transformSlice() {
	a := []A{{Name: "John", Age: 20, C: C{School: "Harvard"}}, {Name: "Doe", Age: 30}}
	b := []B{}
	err := transx.TransformSlice(&a, &b)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonStr, _ := json.Marshal(b)
	fmt.Printf("%+v\n", string(jsonStr))
}

type A struct {
	Name string `transx:"name"`
	Age  int    `transx:"age"`
	C
}

type B struct {
	Nama    string `transx:"name"`
	Umur    int    `transx:"age"`
	Sekolah string `transx:"school"`
}

type C struct {
	School string `transx:"school"`
}
