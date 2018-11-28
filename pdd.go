package pdd

import (
	. "github.com/liunian1004/pdd/context"
	"github.com/liunian1004/pdd/ddk"
	"github.com/liunian1004/pdd/goods"
)

type Pdd struct {
	Context *Context
}

func NewPdd(c *Config) *Pdd {
	return &Pdd{ Context: NewContext(c) }
}

func (p *Pdd) GetDDK() *ddk.DDK {
	return ddk.NewDDKWithContext(p.Context)
}

func (p *Pdd) GetGoodsAPI() *goods.GoodsAPI {
	return goods.NewGoodsAPIWithContext(p.Context)
}
