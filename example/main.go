package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/imloama/gf-nacos"
)

func main(){
	g.Cfg().SetPath("./")
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("hello, nacos!")
	})
	s.Plugin(&gfnacos.GfNacosPlugin{})
	s.Run()
}