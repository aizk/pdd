package pdd

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

func (p *Pdd) GetDDK() *DDK {
	return NewDDKWithContext(p.Context)
}

func (p *Pdd) GetGoodsAPI() *GoodsAPI {
	return NewGoodsAPIWithContext(p.Context)
}
