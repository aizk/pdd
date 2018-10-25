package pdd

import (
	"testing"
	"fmt"
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
	s, err := p.GetThemeList(1, 20)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", s)
}
