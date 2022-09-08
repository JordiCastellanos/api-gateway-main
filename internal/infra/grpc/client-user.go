package grpc

import (
	"sync"
	"template-api-gateway/cmd/config"
	pb "template-api-gateway/internal/infra/ploto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var once sync.Once
var err error

var con *grpc.ClientConn

func NewClientGrpcUser() pb.UserCrudClient {
	once.Do(func() {
		con, err = grpc.Dial(config.GetConfig().Domain.GrpcUser, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic("error al conectarse al grpc")
		}
	})

	return pb.NewUserCrudClient(con)
}
