package api

import (
	"context"
	"doctor/internal/config"
	"doctor/internal/models"
	"doctor/internal/repository"
	"doctor/internal/utils"

	"github.com/google/uuid"
	docpb "github.com/the-spine/spine-protos-go/doctor"
)

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
		Street:     doc.Address.Street,
		City:       doc.Address.City,
		State:      doc.Address.State,
		Country:    doc.Address.Country,
		PostalCode: doc.Address.PostalCode,
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
			Weekday:    utils.ProtoWeekdayToWeekDay(ah.WeekDay),
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

	resDoctors := make([]*docpb.Doctor, 10)

	id := req.DoctorId

	if id != "" {
		if doc, err := repository.GetDoctorByID(id); err == nil {
			d := doc.ToProtoDoctor()
			resDoctors = append(resDoctors, &d)
		}
	} else {
		docs := make([]models.Doctor, 0)
		if err := repository.GetDoctors(&docs); err == nil {
			for _, doc := range docs {
				d := doc.ToProtoDoctor()
				resDoctors = append(resDoctors, &d)
			}
		}
	}
	res := docpb.GetDoctorResponse{
		Doctors: resDoctors,
	}
	return &res, nil
}

func (d *doctorGrpcApi) CheckDoctorAvailability(ctx context.Context, req *docpb.DoctorAvailabilityRequest) (*docpb.DoctorAvailabilityResponse, error) {

	hour := req.CurrentHour

	weekday := utils.ProtoWeekdayToWeekDay(req.WeekDay)

	id := req.DoctorId

	available, err := repository.CheckDoctorAvailability(id, weekday, hour)

	if err != nil {
		return nil, err
	}

	return &docpb.DoctorAvailabilityResponse{
		Available: available,
	}, nil
}

func (d *doctorGrpcApi) UpdateDoctor(ctx context.Context, req *docpb.UpdateDoctorRequest) (*docpb.UpdateDoctorResponse, error) {
	doc := req.GetDoctor()

	id := uuid.New()

	address := models.Address{
		Street:     doc.Address.Street,
		City:       doc.Address.City,
		State:      doc.Address.State,
		Country:    doc.Address.Country,
		PostalCode: doc.Address.PostalCode,
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
			Weekday:    utils.ProtoWeekdayToWeekDay(ah.WeekDay),
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

	res := docpb.UpdateDoctorResponse{
		Success: false,
	}

	if err != nil {
		return &res, err
	}

	res.Success = true

	return &res, nil
}

func (d *doctorGrpcApi) DeleteDoctor(ctx context.Context, req *docpb.DeleteDoctorRequest) (*docpb.DeleteDoctorResponse, error) {
	id := req.DoctorId

	err := repository.DeleteDoctor(id)

	res := docpb.DeleteDoctorResponse{
		Success: false,
	}

	if err != nil {
		return &res, err
	}

	res.Success = true

	return &res, nil
}
