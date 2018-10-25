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
