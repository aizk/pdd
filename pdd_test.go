package pdd

import (
	"testing"
	"fmt"
	"github.com/liunian1004/pdd/context"
)

func TestNewPdd(t *testing.T) {
	p := NewPdd(&context.Config{
		ClientId: "bf0a1f05b9884f3fba9e3beeaaf5c8bb",
		ClientSecret: "84977a0d640cba853fdf1dfac7736cdc9e2df34f",
	})
	d := p.GetDDK()
	fmt.Println(p, d)
	r, err := d.GoodsDetail(1143441478)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v", r)
}

func TestPdd_GetThemeList(t *testing.T) {
	//p := NewPdd(&Config{
	//	ClientId: "",
	//	ClientSecret: "",
	//})
	//d := ddk.NewDDK(p)
	//s, err := d.ThemeListGet(1, 20)
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Logf("%+v", s)
	//
	//SetDebug(true)
	//SetRetryTimes(3)
}
