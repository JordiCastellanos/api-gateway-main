package service

import (
	"context"
	"fmt"
	"template-api-gateway/cmd/config"
	"template-api-gateway/internal/infra"
	pb "template-api-gateway/internal/infra/ploto"

	"template-api-gateway/internal/infra/grpc"
)

type serviceUser struct {
	rest infra.ClientRest
	grpc pb.UserCrudClient
}

func NewServiceUser() serviceUser {

	return serviceUser{
		rest: infra.NewRequest(config.GetConfig().Domain.Google, nil),
		grpc: grpc.NewClientGrpcUser(),
	}
}

func (s *serviceUser) GetData() string {
	response, err := s.rest.Get("/")
	fmt.Println(err)
	return response.Body
}

func (s *serviceUser) GetDataGrpc() *pb.Response {
	response, err := s.grpc.Insert(context.Background(), &pb.User{
		Name: "dasdas",
	})
	fmt.Println(err)
	return response
}
