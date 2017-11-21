package expirations

import "time"

type ExpirationData struct {
	time time.Time
	keys []uint
}

func (e *ExpirationData) ComesBefore(x Expiree) (bool) {
	return e.time.Before(x.getTime())
}

func (e *ExpirationData) IsEqual(x Expiree) (bool) {
	return e.time.Equal(x.getTime())
}

func (e *ExpirationData) Append(x Expiree) {
	e.keys = append(e.keys, x.getKeys()...)
}

func (e *ExpirationData) getKeys() []uint {
	return e.keys
}

func (e *ExpirationData) getTime() time.Time {
	return e.time
}

func NewExpirationData(t time.Time, keys []uint) ExpirationData {
	return ExpirationData{
		time: t,
		keys: keys}
}