package api

import (
	"context"
	"doctor/internal/config"
	"doctor/internal/models"
	"doctor/internal/repository"
	"time"

	"github.com/google/uuid"
	docpb "github.com/the-spine/spine-protos-go/doctor"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func getWeekDay(w docpb.WeekDay) time.Weekday {
	switch w {
	case docpb.WeekDay_MONDAY:
		return time.Monday
	case docpb.WeekDay_TUESDAY:
		return time.Tuesday
	case docpb.WeekDay_WEDNESDAY:
		return time.Wednesday
	case docpb.WeekDay_THURSDAY:
		return time.Thursday
	case docpb.WeekDay_FRIDAY:
		return time.Friday
	case docpb.WeekDay_SATURDAY:
		return time.Saturday
	case docpb.WeekDay_SUNDAY:
		return time.Sunday
	default:
		return time.Monday
	}
}

type doctorGrpcApi struct {
	config config.Config
	docpb.UnimplementedDoctorServiceServer
}

func GetDoctorGrpcApiServiceServer(config config.Config) *doctorGrpcApi {
	return &doctorGrpcApi{
		config: config,
	}
}

func (d *doctorGrpcApi) RegisterDoctor(ctx context.Context, req *docpb.RegisterDoctorRequest) (*docpb.RegisterDoctorResponse, error) {

	doc := req.GetDoctor()

	id := uuid.New()

	address := models.Address{
		Street:     doc.ContactDetails.Address.Street,
		City:       doc.ContactDetails.Address.City,
		State:      doc.ContactDetails.Address.State,
		Country:    doc.ContactDetails.Address.Country,
		PostalCode: doc.ContactDetails.Address.PostalCode,
	}

	contact := models.ContanctDetails{
		Email:            doc.ContactDetails.Email,
		PhoneNumber:      doc.ContactDetails.PhoneNumber,
		Website:          doc.ContactDetails.Website,
		EmergencyContact: doc.ContactDetails.EmergencyContact,
	}

	educations := make([]models.Education, 0)

	for _, edu := range doc.Educations {
		educations = append(educations, models.Education{
			Degree:         edu.Degree,
			University:     edu.University,
			GraduationYear: int(edu.GraduationYear),
		})
	}

	availableHours := make([]models.AvailableHours, 0)

	for _, ah := range doc.AvailableHours {
		availableHours = append(availableHours, models.AvailableHours{
			Weekday:    getWeekDay(ah.WeekDay),
			StartHours: int(ah.StartHour),
			EndHours:   int(ah.EndHour),
		})

	}

	doctor := models.Doctor{
		ID:              id,
		FirstName:       doc.Name.FirstName,
		MiddleName:      doc.Name.MiddleName,
		LastName:        doc.Name.LastName,
		Address:         address,
		ContanctDetails: contact,
		Educations:      educations,
		AvailableHours:  availableHours,
	}

	err := repository.CreateDoctor(&doctor)

	if err != nil {
		return nil, err
	}

	res := docpb.RegisterDoctorResponse{
		DoctorId: id.String(),
	}

	return &res, nil
}
func (d *doctorGrpcApi) GetDoctor(ctx context.Context, req *docpb.GetDoctorRequest) (*docpb.GetDoctorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDoctor not implemented")
}
func (d *doctorGrpcApi) CheckDoctorAvailability(ctx context.Context, req *docpb.DoctorAvailabilityRequest) (*docpb.DoctorAvailabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckDoctorAvailability not implemented")
}
func (d *doctorGrpcApi) UpdateDoctor(ctx context.Context, req *docpb.UpdateDoctorRequest) (*docpb.UpdateDoctorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDoctor not implemented")
}
func (d *doctorGrpcApi) DeleteDoctor(ctx context.Context, req *docpb.DeleteDoctorRequest) (*docpb.DeleteDoctorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDoctor not implemented")
}
