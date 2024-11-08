package main

import (
	"encoding/json"
	"fmt"
	"git.verzth.work/go/transx"
	"github.com/shopspring/decimal"
)

func main() {
	transform()
	transformSlice()
}

func transform() {
	a := A{Name: "John", Age: 20, Point: decimal.NewFromInt(5000), C: C{School: "Harvard"}, ID: D{Number: 1}, IDX: &D{Number: 1}, IDs: []D{{Number: 2}}, IDXs: []*D{{Number: 2}}, List: &X{"a", "b", "c"}, ListX: []X{{"a", "b", "c"}, {"d", "e", "f"}}}
	b := B{}
	err := transx.Transform(a, &b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Single: %+v\n", b)
}

func transformSlice() {
	a := []A{{Name: "John", Age: 20, Point: decimal.NewFromInt(25), C: C{School: "Harvard"}, ID: D{Number: 1}, IDX: &D{Number: 1}, IDs: []D{{Number: 2}}, IDXs: []*D{{Number: 2}}}, {Name: "Doe", Age: 30, ID: D{Number: 2}}}
	b := []B{}
	err := transx.TransformSlice(&a, &b)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonStr, _ := json.Marshal(b)
	fmt.Printf("List: %+v\n", string(jsonStr))
}

type X []string

type A struct {
	Name  string          `transx:"name"`
	Age   int             `transx:"age"`
	Point decimal.Decimal `transx:"point"`
	C
	ID    D    `transx:"id"`
	IDX   *D   `transx:"idx"`
	IDs   []D  `transx:"ids"`
	IDXs  []*D `transx:"idxs"`
	List  *X   `transx:"list"`
	ListX []X  `transx:"listx"`
}

type B struct {
	Nama    string          `transx:"name"`
	Umur    int             `transx:"age"`
	Point   decimal.Decimal `transx:"point"`
	Sekolah string          `transx:"school"`
	ID      E               `transx:"id,deep"`
	IDX     *D              `transx:"idx,deep"`
	IDs     []E             `transx:"ids,deep"`
	IDXs    []*E            `transx:"idxs,deep"`
	List    *X              `transx:"list,deep"`
	ListX   []X             `transx:"listx,deep"`
}

type C struct {
	School string `transx:"school"`
}

type D struct {
	Number int `transx:"number"`
}

type E struct {
	Nomor int `transx:"number"`
}
