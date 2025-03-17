package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	dataModels "github.com/nambuitechx/go-metadata/models/data"
	servicesRepositories "github.com/nambuitechx/go-metadata/repositories/services"
	dataRepositories "github.com/nambuitechx/go-metadata/repositories/data"
)

type DatabaseSchemaEntityService struct {
	DBServiceEntityRepository *servicesRepositories.DBServiceEntityRepository
	DatabaseEntityRepository *dataRepositories.DatabaseEntityRepository
	DatabaseSchemaEntityRepository *dataRepositories.DatabaseSchemaEntityRepository
}

func NewDatabaseSchemaEntityService(
	dbserviceEntityRepository *servicesRepositories.DBServiceEntityRepository,
	databaseEntityRepository *dataRepositories.DatabaseEntityRepository,
	databaseSchemaEntityRepository *dataRepositories.DatabaseSchemaEntityRepository,
) *DatabaseSchemaEntityService {
	return &DatabaseSchemaEntityService{
		DBServiceEntityRepository: dbserviceEntityRepository,
		DatabaseEntityRepository: databaseEntityRepository,
		DatabaseSchemaEntityRepository: databaseSchemaEntityRepository,
	}
}

func (s *DatabaseSchemaEntityService) Health() string {
	return "Database schema service is available"
}

func (s *DatabaseSchemaEntityService) GetAllDatabaseSchemaEntities(limit int, offset int) ([]dataModels.DatabaseSchemaEntity, error) {
	databaseSchemaEntity, err := s.DatabaseSchemaEntityRepository.SelectDatabaseSchemaEntities(limit, offset)
	return databaseSchemaEntity, err
}

func (s *DatabaseSchemaEntityService) GetDatabaseSchemaEntityById(id string) (*dataModels.DatabaseSchemaEntity, error) {
	databaseSchemaEntity, err := s.DatabaseSchemaEntityRepository.SelectDatabaseSchemaEntityById(id)
	return databaseSchemaEntity, err
}

func (s *DatabaseSchemaEntityService) GetDatabaseSchemaEntityByFqn(fqn string) (*dataModels.DatabaseSchemaEntity, error) {
	databaseSchemaEntity, err := s.DatabaseSchemaEntityRepository.SelectDatabaseSchemaEntityByFqn(fqn)
	return databaseSchemaEntity, err
}

func (s *DatabaseSchemaEntityService) CreateDatabaseSchemaEntity(payload *dataModels.CreateDatabaseSchemaEntityPayload) (*dataModels.DatabaseSchemaEntity, error) {
	id := uuid.NewString()
	now := time.Now().Unix()

	// Split dbservice and database
	arr := strings.Split(payload.Database, ".")

	if len(arr) != 2 {
		return nil, errors.New("invalid database field")
	}

	// Get dbservice
	dbservice, err := s.DBServiceEntityRepository.SelectDBServiceEntityByFqn(arr[0])

	if err != nil {
		return nil, err
	}

	dbserviceEntityRef := dbservice.Json.ToEntityReference()

	// Get database
	database, err := s.DatabaseEntityRepository.SelectDatabaseEntityByFqn(fmt.Sprintf("%v.%v", arr[0], arr[1]))

	if err != nil {
		return nil, err
	}

	databaseEntityRef := database.Json.ToEntityReference()

	// Populate database schema
	databaseSchema := &dataModels.DatabaseSchema{
		ID: id,
		Name: payload.Name,
		FullyQualifiedName: fmt.Sprintf("%v.%v", payload.Database, payload.Name),
		DisplayName: payload.DisplayName,
		Description: payload.Description,
		ServiceType: dbservice.ServiceType,
		Service: dbserviceEntityRef,
		Database: databaseEntityRef,
		Deleted: false,
	}

	entity := &dataModels.DatabaseSchemaEntity{
		ID: id,
		Name: payload.Name,
		Json: databaseSchema,
		UpdatedAt: now,
		Deleted: false,
	}

	databaseSchemaEntity, err := s.DatabaseSchemaEntityRepository.InsertDatabaseSchemaEntity(entity)
	return databaseSchemaEntity, err
}

func (s *DatabaseSchemaEntityService) DeleteDatabaseSchemaEntityById(id string) error {
	err := s.DatabaseSchemaEntityRepository.DeleteDatabaseSchemaEntityById(id)
	return err
}

func (s *DatabaseSchemaEntityService) DeleteDatabaseSchemaEntityByFqn(fqn string) error {
	err := s.DatabaseSchemaEntityRepository.DeleteDatabaseSchemaEntityByFqn(fqn)
	return err
}
