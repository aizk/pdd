package pdd

type Context struct {
	ClientId     string
	ClientSecret string
	RetryTimes   int
	Debug        bool
}

func NewContext(cfg *Config) *Context {
	return &Context{
		ClientId:     cfg.ClientId,
		ClientSecret: cfg.ClientSecret,
		RetryTimes:   cfg.RetryTimes,
		Debug:        cfg.Debug,
	}
}
