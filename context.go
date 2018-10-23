package pdd

type Context struct {
	ClientId     string
	ClientSecret string
}

func NewContext(c *Config) *Context {
	return &Context{
		ClientId: c.ClientId,
		ClientSecret: c.ClientSecret,
	}
}