package model

type HelloDao interface {
}

type helloDao struct {
}

func NewHelloDao() *helloDao {
	return &helloDao{}
}
