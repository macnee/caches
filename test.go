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
	fmt.Printf("%v\n", v)


	//more intense memery
	first := expirations.NewExpirationData(time.Now().Add(time.Hour * 1), []uint{1})
	second := expirations.NewExpirationData(time.Now().Add(time.Hour * 5), []uint{2})
	third := expirations.NewExpirationData(time.Now().Add(time.Hour * 2), []uint{3})
	fourth := expirations.NewExpirationData(time.Now().Add(time.Hour* 7), []uint{4})
	fifth := expirations.NewExpirationData(time.Now().Add(time.Hour * 5), []uint{5})

	a.Push(&first)
	a.Push(&second)
	a.Push(&third)
	a.Push(&fourth)
	a.Push(&fifth)

	fmt.Printf("%v", a.Pop())
	fmt.Printf("%v", a.Pop())
	fmt.Printf("%v", a.Pop())
	fmt.Printf("%v", a.Pop())

}
