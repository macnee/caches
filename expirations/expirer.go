package expirations

import "time"

//structs
type expirationData struct {
	time time.Time
	keys []uint
}

//interfaces
type expiree interface {
	NewExpiree() interface{}
	getKeys() []uint
	getTime() time.Time
}


//implementations


