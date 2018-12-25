package ddk

import (
	"testing"
	"github.com/liunian1004/pdd/context"
	"fmt"
)

func TestPdd_GenerateRedPackagePromUrl(t *testing.T) {

	d := NewDDKWithContext(&context.Context{
		ClientId: "bf0a1f05b9884f3fba9e3beeaaf5c8bb",
		ClientSecret: "84977a0d640cba853fdf1dfac7736cdc9e2df34f",
	})
	//s, err := d.GetThemeList(1, 20)
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Logf("%+v", s)

	//params := util.NewParams()
	//params.Set("custom_parameters", "80008_88_redpackage")
	//params.Set("generate_short_url", true)
	//s1, err := d.RPPromUrlGenerate([]string{"1891891_26364337"}, true, params)
	//if err != nil {
	//	t.Error(err)
	//}
	//bytes, _ := json.Marshal(s1[0])
	//t.Logf("%+v", *s1[0])
	//t.Log(string(bytes))

	fmt.Println(d.GoodsDetail(937925432))
}
