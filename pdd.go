package pdd

type Pdd struct {
	Context *Context
}

type Config struct {
	ClientId     string
	ClientSecret string
}

func NewPdd(c *Config) *Pdd {
	return &Pdd{ Context: NewContext(c) }
}

// the basic call method
func (p *Pdd) Call(params Params) (r []byte, err error) {
	prs := NewParams()
	prs.SetParams(params)
	prs.Sign(p)
	r, err = GetBytesPost(params.GetQuery())
	return
}