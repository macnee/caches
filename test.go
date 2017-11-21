package main

import (
	"fmt"
	"github.com/macnee/caches/expirations"
	"time"
)

func main() {
	fmt.Printf("starting tests for my memes\n")
	e := expirations.NewExpirationData(time.Now().Add(time.Hour), []uint{1})
	h := expirations.NewExpirationHeap()

	var b expirations.Expiree
	var a expirations.Expirer
	b = &e
	a = &h

	a.Push(b)
	v := a.Pop()
	fmt.Printf("%v", v)


	//more intense memery
}
