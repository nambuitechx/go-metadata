package services

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	dataModels "github.com/nambuitechx/go-metadata/models/data"
	servicesRepositories "github.com/nambuitechx/go-metadata/repositories/services"
	dataRepositories "github.com/nambuitechx/go-metadata/repositories/data"
)

type DatabaseEntityService struct {
	DBServiceEntityRepository *servicesRepositories.DBServiceEntityRepository
	DatabaseEntityRepository *dataRepositories.DatabaseEntityRepository
}

func NewDatabaseEntityService(
	dbserviceEntityRepository *servicesRepositories.DBServiceEntityRepository,
	databaseEntityRepository *dataRepositories.DatabaseEntityRepository,
) *DatabaseEntityService {
	return &DatabaseEntityService{
		DBServiceEntityRepository: dbserviceEntityRepository,
		DatabaseEntityRepository: databaseEntityRepository,
	}
}

func (s *DatabaseEntityService) Health() string {
	return "Database service is available"
}

func (s *DatabaseEntityService) GetAllDatabaseEntities(limit int, offset int) ([]dataModels.DatabaseEntity, error) {
	databaseEntity, err := s.DatabaseEntityRepository.SelectDatabaseEntities(limit, offset)
	return databaseEntity, err
}

func (s *DatabaseEntityService) GetDatabaseEntityById(id string) (*dataModels.DatabaseEntity, error) {
	databaseEntity, err := s.DatabaseEntityRepository.SelectDatabaseEntityById(id)
	return databaseEntity, err
}

func (s *DatabaseEntityService) GetDatabaseEntityByFqn(fqn string) (*dataModels.DatabaseEntity, error) {
	databaseEntity, err := s.DatabaseEntityRepository.SelectDatabaseEntityByFqn(fqn)
	return databaseEntity, err
}

func (s *DatabaseEntityService) CreateDatabaseEntity(payload *dataModels.CreateDatabaseEntityPayload) (*dataModels.DatabaseEntity, error) {
	id := uuid.NewString()
	now := time.Now().Unix()

	// Get dbservice
	dbservice, err := s.DBServiceEntityRepository.SelectDBServiceEntityByFqn(payload.Service)

	if err != nil {
		return nil, err
	}

	dbserviceEntityRef := dbservice.Json.ToEntityReference()

	// Populate database
	database := &dataModels.Database{
		ID: id,
		Name: payload.Name,
		FullyQualifiedName: fmt.Sprintf("%v.%v", payload.Service, payload.Name),
		DisplayName: payload.DisplayName,
		Description: payload.Description,
		ServiceType: dbservice.ServiceType,
		Service: dbserviceEntityRef,
		Deleted: false,
	}

	entity := &dataModels.DatabaseEntity{
		ID: id,
		Name: payload.Name,
		Json: database,
		UpdatedAt: now,
		Deleted: false,
	}

	databaseEntity, err := s.DatabaseEntityRepository.InsertDatabaseEntity(entity)
	return databaseEntity, err
}

func (s *DatabaseEntityService) DeleteDatabaseEntityById(id string) error {
	err := s.DatabaseEntityRepository.DeleteDatabaseEntityById(id)
	return err
}

func (s *DatabaseEntityService) DeleteDatabaseEntityByFqn(fqn string) error {
	err := s.DatabaseEntityRepository.DeleteDatabaseEntityByFqn(fqn)
	return err
}
