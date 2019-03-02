package pdd

type Pdd struct {
	Context *Context
}

type Config struct {
	ClientId     string
	ClientSecret string
	RetryTimes   int
	Debug        bool
}

func NewPdd(c *Config) *Pdd {
	return &Pdd{
		Context: NewContext(c),
	}
}

func (p *Pdd) GetDDK() *DDK {
	return newDDKWithContext(p.Context)
}

func (p *Pdd) GetGoodsAPI() *GoodsAPI {
	return newGoodsAPIWithContext(p.Context)
}
