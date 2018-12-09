package main

import (
	"fmt"
	"github.com/dy-dayan/ap-tcp/handler"
	"github.com/dy-dayan/ap-tcp/idl"
	"github.com/dy-gopkg/kit"
)

func main() {

	kit.Init()

	h := handler.NewHandler()

	access.RegisterAccessHandler(kit.Server(), h)

	go kit.Run()

	//TODO: need a delay to start?
	h.Start()

	fmt.Println("aaa")
}
