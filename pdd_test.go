package pdd

import (
	"testing"
	"fmt"
	"github.com/liunian1004/pdd/ddk"
)

func TestNewPdd(t *testing.T) {
	p := NewPdd(&Config{})
	fmt.Println(p)
}

func TestPdd_GetThemeList(t *testing.T) {
	p := NewPdd(&Config{
		ClientId: "",
		ClientSecret: "",
	})
	d := ddk.NewDDK(p)
	s, err := d.ThemeListGet(1, 20)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", s)

	SetDebug(true)
	SetRetryTimes(3)
}
