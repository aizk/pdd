package context

type Config struct {
	ClientId     string
	ClientSecret string
}

type Context struct {
	ClientId     string
	ClientSecret string
	Retry        int
	Debug        bool
}

func NewContext(c *Config) *Context {
	return &Context{
		ClientId:     c.ClientId,
		ClientSecret: c.ClientSecret,
	}
}
