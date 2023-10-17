package api

import (
	"context"
	"doctor/internal/config"

	"github.com/the-spine/spine-protos-go/doctor"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type doctorGrpcApi struct {
	config config.Config
	doctor.UnimplementedDoctorServiceServer
}

func GetDoctorGrpcApiServiceServer(config config.Config) *doctorGrpcApi {
	return &doctorGrpcApi{
		config: config,
	}
}

func (d *doctorGrpcApi) RegisterDoctor(context.Context, *doctor.RegisterDoctorRequest) (*doctor.RegisterDoctorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDoctor not implemented")
}
func (d *doctorGrpcApi) GetDoctor(context.Context, *doctor.GetDoctorRequest) (*doctor.GetDoctorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDoctor not implemented")
}
func (d *doctorGrpcApi) CheckDoctorAvailability(context.Context, *doctor.DoctorAvailabilityRequest) (*doctor.DoctorAvailabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckDoctorAvailability not implemented")
}
func (d *doctorGrpcApi) UpdateDoctor(context.Context, *doctor.UpdateDoctorRequest) (*doctor.UpdateDoctorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDoctor not implemented")
}
func (d *doctorGrpcApi) DeleteDoctor(context.Context, *doctor.DeleteDoctorRequest) (*doctor.DeleteDoctorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDoctor not implemented")
}
