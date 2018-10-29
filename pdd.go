package pdd

import (
	. "github.com/liunian1004/pdd/context"
	"github.com/liunian1004/pdd/ddk"
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