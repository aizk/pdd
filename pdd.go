package pdd

import (
	. "github.com/liunian1004/pdd/context"
	"github.com/liunian1004/pdd/ddk"
	"github.com/liunian1004/pdd/goods"
)

type Pdd struct {
	Context *Context
}

type Config struct {
	ClientId     string
	ClientSecret string
}

func NewPdd(c *Config) *Pdd {
	return &Pdd{ Context: NewContext(c.ClientId, c.ClientSecret) }
}

func (p *Pdd) GetDDK() *ddk.DDK {
	return ddk.NewDDKWithContext(p.Context)
}

func (p *Pdd) GetGoodsAPI() *goods.GoodsAPI {
	return goods.NewGoodsAPIWithContext(p.Context)
}
