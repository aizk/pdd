package ddk

import (
	"testing"
	"github.com/liunian1004/pdd"
)

func TestPdd_GenerateRedPackagePromUrl(t *testing.T) {
	p := pdd.NewPdd(&pdd.Config{
		ClientId: "",
		ClientSecret: "",
	})
	d := NewDDK(p)
	s, err := d.ThemeListGet(1, 20)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", s)

	params := pdd.NewParams()
	params.Set("custom_parameters", "80008_88_redpackage")
	params.Set("generate_short_url", true)
	s1, err := d.RPPromUrlGenerate([]string{"1891891_26364337"}, true, params)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", *s1[0])
}
