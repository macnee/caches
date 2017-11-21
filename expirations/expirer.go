package expirations

//interfaces
type Expirer interface {
	Push(x Expiree)
	Pop() []uint
}