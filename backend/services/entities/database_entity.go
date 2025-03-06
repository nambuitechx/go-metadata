package services

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nambuitechx/go-metadata/models/entities"
	"github.com/nambuitechx/go-metadata/repositories/entities"
)

type DatabaseEntityService struct {
	DBServiceEntityRepository *repositories.DBServiceEntityRepository
	DatabaseEntityRepository *repositories.DatabaseEntityRepository
}

func NewDatabaseEntityService(
	dbserviceEntityRepository *repositories.DBServiceEntityRepository,
	databaseEntityRepository *repositories.DatabaseEntityRepository,
) *DatabaseEntityService {
	return &DatabaseEntityService{
		DBServiceEntityRepository: dbserviceEntityRepository,
		DatabaseEntityRepository: databaseEntityRepository,
	}
}

func (s *DatabaseEntityService) Health() string {
	return "Database service is available"
}

func (s *DatabaseEntityService) GetAllDatabaseEntities(limit int, offset int) ([]models.DatabaseEntity, error) {
	databaseEntity, err := s.DatabaseEntityRepository.SelectDatabaseEntities(limit, offset)
	return databaseEntity, err
}

func (s *DatabaseEntityService) GetDatabaseEntityById(id string) (*models.DatabaseEntity, error) {
	databaseEntity, err := s.DatabaseEntityRepository.SelectDatabaseEntityById(id)
	return databaseEntity, err
}

func (s *DatabaseEntityService) GetDatabaseEntityByFqn(fqn string) (*models.DatabaseEntity, error) {
	databaseEntity, err := s.DatabaseEntityRepository.SelectDatabaseEntityByFqn(fqn)
	return databaseEntity, err
}

func (s *DatabaseEntityService) CreateDatabaseEntity(payload *models.CreateDatabaseEntityPayload) (*models.DatabaseEntity, error) {
	id := uuid.NewString()
	now := time.Now().Unix()

	// Get dbservice
	dbservice, err := s.DBServiceEntityRepository.SelectDBServiceEntityByFqn(payload.Service)

	if err != nil {
		return nil, err
	}

	dbserviceEntityRef := dbservice.Json.ToEntityReference()

	// Populate database
	database := &models.Database{
		ID: id,
		Name: payload.Name,
		FullyQualifiedName: fmt.Sprintf("%v.%v", payload.Service, payload.Name),
		DisplayName: payload.DisplayName,
		Description: payload.Description,
		ServiceType: dbservice.ServiceType,
		Service: dbserviceEntityRef,
		Deleted: false,
	}

	entity := &models.DatabaseEntity{
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
