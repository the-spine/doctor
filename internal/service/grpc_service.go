package service

import (
	"doctor/internal/api"
	"doctor/internal/config"
	"fmt"
	"net"

	"github.com/the-spine/spine-protos-go/doctor"
	"google.golang.org/grpc"
)

func StartGrpcService(config *config.Config) (*grpc.Server, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Api.Host, config.Api.Port))
	if err != nil {
		return nil, err
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	doctor.RegisterDoctorServiceServer(grpcServer, api.GetDoctorGrpcApiServiceServer(*config))

	grpcServer.Serve(lis)

	return grpcServer, nil
}
