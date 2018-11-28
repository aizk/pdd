package pdd

import (
	"testing"
	"fmt"
	"github.com/liunian1004/pdd/context"
)

func TestNewPdd(t *testing.T) {
	p := NewPdd(&context.Config{
		ClientId: "",
		ClientSecret: "",
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
	//s, err := d.GetThemeList(1, 20)
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Logf("%+v", s)
	//
	//SetDebug(true)
	//SetRetryTimes(3)
}
