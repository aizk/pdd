package context

type Context struct {
	ClientId     string
	ClientSecret string
	Retry        int
	Debug        bool
}

func NewContext(clientId, clientSecet string) *Context {
	return &Context{
		ClientId:     clientId,
		ClientSecret: clientSecet,
	}
}
