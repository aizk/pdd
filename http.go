package pdd

import (
	"github.com/parnurzeal/gorequest"
	"sync"
	"fmt"
)

const EndPoint = "https://gw-api.pinduoduo.com/api/router"

var clients *sync.Pool

func init() {
	clients = &sync.Pool{
		New: func() interface{} {
			return gorequest.New()
		},
	}
}

func GetStructPost(data interface{}, r interface{}) (err error) {
	client := clients.Get().(*gorequest.SuperAgent)
	_, r, errors := client.Post(EndPoint).
		Type("json").
		Query(data).
		EndStruct(r)
	// return back
	clients.Put(client)
	for _, e := range errors {
		if e != nil {
			return e
		}
	}
	return
}

func GetStringPost(query string) (s string, err error) {
	client := clients.Get().(*gorequest.SuperAgent)
	_, s, errors := client.Post(EndPoint).
		Type("json").
		Query(query).
		End()
	// push back
	clients.Put(client)
	if IsBadPddRequest([]byte(s)) {
		err = fmt.Errorf("%s", s)
		return
	}

	for _, e := range errors {
		if e != nil {
			err = e
			return
		}
	}
	return
}

func GetBytesPost(query string) (b []byte, err error) {
	client := clients.Get().(*gorequest.SuperAgent)
	_, b, errors := client.Post(EndPoint).
		Type("json").
		Query(query).
		EndBytes()
	// push back
	clients.Put(client)
	if IsBadPddRequest(b) {
		err = fmt.Errorf("%s", b)
		return
	}

	for _, e := range errors {
		if e != nil {
			err = e
			return
		}
	}
	return
}

func IsBadPddRequest(body []byte) bool {
	return string(body[2:16]) == "error_response"
}