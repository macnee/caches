package cache

import (
	"github.com/snuffalo/caches/expirations"
)

var InsertStream = make(chan Insert)
var GetStream = make(chan Get)

type TemporalCacher interface {
	Insert(s string, e expirations.Expiree) uint
	Get(i uint) string
	Run()
}

type Insert struct {
	s string
	e expirations.Expiree
	response chan uint
}

type Get struct {
	i uint
	response chan string
}