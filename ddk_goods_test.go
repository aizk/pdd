package pdd

import (
	"testing"
	"fmt"
)

func TestGoodsAPI_GoodsCatGet(t *testing.T) {
	g := NewGoodsAPIWithContext(&Context{
		ClientId: ClientId,
		ClientSecret: ClientSecret,
	})
	r, err := g.GoodsCatGet(0)
	fmt.Println(err)
	fmt.Printf("%v", r)
	for _, value := range r {
		fmt.Printf("%#v\n", value)
	}
}
