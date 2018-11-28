package goods

import (
	"testing"
	"github.com/liunian1004/pdd/context"
	"fmt"
	"github.com/liunian1004/pdd/util"
)

func TestGoodsAPI_GoodsCatGet(t *testing.T) {
	g := NewGoodsAPIWithContext(&context.Context{
		ClientId: util.ClientId,
		ClientSecret: util.ClientSecret,
	})
	r, err := g.GoodsCatGet(0)
	fmt.Println(err)
	fmt.Printf("%v", r)
	for _, value := range r {
		fmt.Printf("%#v\n", value)
	}
}
