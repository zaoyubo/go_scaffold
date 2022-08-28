package grpc

import (
	model "template/internal/model/hello"
	service "template/internal/service/hello"
)

func Init() *Api {
	dao := model.NewHelloDao()
	helloService := service.NewHelloService(dao)
	api := NewGrpcAPi(helloService)
	return api
}
