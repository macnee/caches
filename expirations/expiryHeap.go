package expirations

import "time"

//structs
type simpleExpirer struct {
	tree []expiree
}

//interfaces
type expirer interface {
	NewExpirer() interface{}
	InsertByTime(t time.Time) (bool)
	InsertByDuration(d time.Duration) (bool)
	Pop() []uint
}

//implementations