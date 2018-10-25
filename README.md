# README

本项目为 `go` 语言实现的拼多多开放平台 SDK，调用方式简单粗暴。

`go get github.com/liunian1004/pdd`

```go
import github.com/liunian1004/pdd

p := pdd.NewPdd(&pdd.Config{
    ClientId: "your client id",
    ClientSecret: "your client secret",
})

// get theme list
r, err := p.GetThemeList(1, 20)
```
## Todo

- [] 实现所有接口
