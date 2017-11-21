package expirations

import "time"

type Expiree interface {
	ComesBefore(x Expiree) bool
	IsEqual(x Expiree) bool
	Append(x Expiree)
	getKeys() []uint
	getTime() time.Time
}
