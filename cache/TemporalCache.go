package cache

import (
	"github.com/snuffalo/caches/expirations"
)

type TemporalCache struct {
	data map[uint]string
	eheap expirations.Expirer
	nextKey uint
}

func (c *TemporalCache) Run() {
	go func() {
		for request := range InsertStream {
			request.response <- c.Insert(request.s, request.e)
		}
	}()
}
func (c *TemporalCache) Insert(s string, e expirations.Expiree) uint {
	c.data[c.nextKey] = s
	defer c.bumpKey()

	c.eheap.Push(e)
	return c.nextKey
}
func (c *TemporalCache) Get(i uint) string {
	if val, ok := c.data[i]; ok {
		return val
	} else {
		return ""
	}
}

func (c *TemporalCache) bumpKey() {
	c.nextKey = c.nextKey + 1
}

func NewTemporalCache() TemporalCache {

}