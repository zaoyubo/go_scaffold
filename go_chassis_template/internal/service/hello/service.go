package hello

import (
	"context"

	model "template/internal/model/hello"
	pb "template/protocol/grpc/go"
)

type helloService struct {
	helloDao model.HelloDao
}

func NewHelloService(dao model.HelloDao) *helloService {
	return &helloService{
		helloDao: dao,
	}
}

func (h *helloService) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloRes, error) {
	s := "hello"
	requestId := req.GetHeader().GetRequestId()
	var retCode int32 = 0
	message := "success"
	return &pb.HelloRes{
		Header: &pb.RespHeader{
			RequestId: &requestId,
			Retcode:   &retCode,
			Message:   &message,
		},
		HelloResStr: &s,
	}, nil
}
