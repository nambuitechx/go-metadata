package services

import (
	"time"

	"github.com/google/uuid"
	baseModels "github.com/nambuitechx/go-metadata/models/base"
	servicesModels "github.com/nambuitechx/go-metadata/models/services"
	servicesRepositories "github.com/nambuitechx/go-metadata/repositories/services"
)

type DBServiceEntityService struct {
	DBServiceEntityRepository *servicesRepositories.DBServiceEntityRepository
}

func NewDBServiceEntityService(dbserviceEntityRepository *servicesRepositories.DBServiceEntityRepository) *DBServiceEntityService {
	return &DBServiceEntityService{ DBServiceEntityRepository: dbserviceEntityRepository }
}

func (s *DBServiceEntityService) Health() string {
	return "DBService service is available"
}

func (s *DBServiceEntityService) GetAllDBServiceEntities(limit int, offset int) ([]servicesModels.DBServiceEntity, error) {
	dbserviceEntity, err := s.DBServiceEntityRepository.SelectDBServiceEntities(limit, offset)
	return dbserviceEntity, err
}

func (s *DBServiceEntityService) GetCountDBServiceEntities() (*baseModels.EntityTotal, error) {
	entityTotal, err := s.DBServiceEntityRepository.SelectCountDBServiceEntities()
	return entityTotal, err
}

func (s *DBServiceEntityService) GetDBServiceEntityById(id string) (*servicesModels.DBServiceEntity, error) {
	dbserviceEntity, err := s.DBServiceEntityRepository.SelectDBServiceEntityById(id)
	return dbserviceEntity, err
}

func (s *DBServiceEntityService) GetDBServiceEntityByFqn(fqn string) (*servicesModels.DBServiceEntity, error) {
	dbserviceEntity, err := s.DBServiceEntityRepository.SelectDBServiceEntityByFqn(fqn)
	return dbserviceEntity, err
}

func (s *DBServiceEntityService) CreateDBServiceEntity(payload *servicesModels.CreateDBServiceEntityPayload) (*servicesModels.DBServiceEntity, error) {
	id := uuid.NewString()
	now := time.Now().Unix()

	dbservice := &servicesModels.DBService{
		ID: id,
		Name: payload.Name,
		FullyQualifiedName: payload.Name,
		DisplayName: payload.DisplayName,
		Description: payload.Description,
		ServiceType: payload.ServiceType,
		Connection: payload.Connection,
		Deleted: false,
	}

	entity := &servicesModels.DBServiceEntity{
		ID: id,
		Name: payload.Name,
		ServiceType: payload.ServiceType,
		Json: dbservice,
		UpdatedAt: now,
		Deleted: false,
	}

	dbserviceEntity, err := s.DBServiceEntityRepository.InsertDBServiceEntity(entity)
	return dbserviceEntity, err
}

func (s *DBServiceEntityService) CreateOrUpdateDBServiceEntity(payload *servicesModels.CreateDBServiceEntityPayload) (*servicesModels.DBServiceEntity, error) {
	exist, err := s.DBServiceEntityRepository.SelectDBServiceEntityByFqn(payload.Name)

	if err == nil {
		updated, err := s.DBServiceEntityRepository.UpdateDBServiceEntity(exist)
		return updated, err
	}

	id := uuid.NewString()
	now := time.Now().Unix()

	dbservice := &servicesModels.DBService{
		ID: id,
		Name: payload.Name,
		FullyQualifiedName: payload.Name,
		DisplayName: payload.DisplayName,
		Description: payload.Description,
		ServiceType: payload.ServiceType,
		Connection: payload.Connection,
		Deleted: false,
	}

	entity := &servicesModels.DBServiceEntity{
		ID: id,
		Name: payload.Name,
		ServiceType: payload.ServiceType,
		Json: dbservice,
		UpdatedAt: now,
		Deleted: false,
	}

	dbserviceEntity, err := s.DBServiceEntityRepository.InsertDBServiceEntity(entity)
	return dbserviceEntity, err
}

func (s *DBServiceEntityService) DeleteDBServiceEntityById(id string) error {
	err := s.DBServiceEntityRepository.DeleteDBServiceEntityById(id)
	return err
}

func (s *DBServiceEntityService) DeleteDBServiceEntityByFqn(fqn string) error {
	err := s.DBServiceEntityRepository.DeleteDBServiceEntityByFqn(fqn)
	return err
}
