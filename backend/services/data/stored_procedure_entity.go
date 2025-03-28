package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	baseModels "github.com/nambuitechx/go-metadata/models/base"
	dataModels "github.com/nambuitechx/go-metadata/models/data"
	servicesRepositories "github.com/nambuitechx/go-metadata/repositories/services"
	dataRepositories "github.com/nambuitechx/go-metadata/repositories/data"
)

type StoredProcedureEntityService struct {
	DBServiceEntityRepository *servicesRepositories.DBServiceEntityRepository
	DatabaseEntityRepository *dataRepositories.DatabaseEntityRepository
	DatabaseSchemaEntityRepository *dataRepositories.DatabaseSchemaEntityRepository
	StoredProcedureEntityRepository *dataRepositories.StoredProcedureEntityRepository
}

func NewStoredProcedureEntityService(
	dbserviceEntityRepository *servicesRepositories.DBServiceEntityRepository,
	databaseEntityRepository *dataRepositories.DatabaseEntityRepository,
	databaseSchemaEntityRepository *dataRepositories.DatabaseSchemaEntityRepository,
	storedProcedureEntityRepository *dataRepositories.StoredProcedureEntityRepository,
) *StoredProcedureEntityService {
	return &StoredProcedureEntityService{
		DBServiceEntityRepository: dbserviceEntityRepository,
		DatabaseEntityRepository: databaseEntityRepository,
		DatabaseSchemaEntityRepository: databaseSchemaEntityRepository,
		StoredProcedureEntityRepository: storedProcedureEntityRepository,
	}
}

func (s *StoredProcedureEntityService) Health() string {
	return "Stored procedure service is available"
}

func (s *StoredProcedureEntityService) GetAllStoredProcedureEntities(databaseSchema string, limit int, offset int) ([]dataModels.StoredProcedureEntity, error) {
	storedProcedureEntity, err := s.StoredProcedureEntityRepository.SelectStoredProcedureEntities(databaseSchema, limit, offset)
	return storedProcedureEntity, err
}

func (s *StoredProcedureEntityService) GetCountStoredProcedureEntities(databaseSchema string) (*baseModels.EntityTotal, error) {
	entityTotal, err := s.StoredProcedureEntityRepository.SelectCountStoredProcedureEntities(databaseSchema)
	return entityTotal, err
}

func (s *StoredProcedureEntityService) GetStoredProcedureEntityById(id string) (*dataModels.StoredProcedureEntity, error) {
	storedProcedureEntity, err := s.StoredProcedureEntityRepository.SelectStoredProcedureEntityById(id)
	return storedProcedureEntity, err
}

func (s *StoredProcedureEntityService) GetStoredProcedureEntityByFqn(fqn string) (*dataModels.StoredProcedureEntity, error) {
	storedProcedureEntity, err := s.StoredProcedureEntityRepository.SelectStoredProcedureEntityByFqn(fqn)
	return storedProcedureEntity, err
}

func (s *StoredProcedureEntityService) CreateStoredProcedureEntity(payload *dataModels.CreateStoredProcedureEntityPayload) (*dataModels.StoredProcedureEntity, error) {
	id := uuid.NewString()
	now := time.Now().Unix()

	// Split dbservice and database
	arr := strings.Split(payload.DatabaseSchema, ".")

	if len(arr) != 3 {
		return nil, errors.New("invalid database schema field")
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

	// Get database schema
	databaseSchema, err := s.DatabaseSchemaEntityRepository.SelectDatabaseSchemaEntityByFqn(fmt.Sprintf("%v.%v.%v", arr[0], arr[1], arr[2]))

	if err != nil {
		return nil, err
	}

	databaseSchemaEntityRef := databaseSchema.Json.ToEntityReference()

	// Populate stored procedure
	storedProcedure := &dataModels.StoredProcedure{
		ID: id,
		Name: payload.Name,
		FullyQualifiedName: fmt.Sprintf("%v.%v", payload.DatabaseSchema, payload.Name),
		DisplayName: payload.DisplayName,
		Description: payload.Description,
		StoredProcedureCode: payload.StoredProcedureCode,
		StoredProcedureType: payload.StoredProcedureType,
		ServiceType: dbservice.ServiceType,
		Service: dbserviceEntityRef,
		Database: databaseEntityRef,
		DatabaseSchema: databaseSchemaEntityRef,
		Deleted: false,
	}

	entity := &dataModels.StoredProcedureEntity{
		ID: id,
		Name: payload.Name,
		Json: storedProcedure,
		UpdatedAt: now,
		Deleted: false,
	}

	storedProcedureEntity, err := s.StoredProcedureEntityRepository.InsertStoredProcedureEntity(entity)
	return storedProcedureEntity, err
}

func (s *StoredProcedureEntityService) CreateOrUpdateStoredProcedureEntity(payload *dataModels.CreateStoredProcedureEntityPayload) (*dataModels.StoredProcedureEntity, error) {
	exist, err := s.StoredProcedureEntityRepository.SelectStoredProcedureEntityByFqn(fmt.Sprintf("%v.%v", payload.DatabaseSchema, payload.Name))

	if err == nil {
		updated, err := s.StoredProcedureEntityRepository.UpdateStoredProcedureEntity(exist)
		return updated, err
	}

	id := uuid.NewString()
	now := time.Now().Unix()

	// Split dbservice and database
	arr := strings.Split(payload.DatabaseSchema, ".")

	if len(arr) != 3 {
		return nil, errors.New("invalid database schema field")
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

	// Get database schema
	databaseSchema, err := s.DatabaseSchemaEntityRepository.SelectDatabaseSchemaEntityByFqn(fmt.Sprintf("%v.%v.%v", arr[0], arr[1], arr[2]))

	if err != nil {
		return nil, err
	}

	databaseSchemaEntityRef := databaseSchema.Json.ToEntityReference()

	// Populate stored procedure
	storedProcedure := &dataModels.StoredProcedure{
		ID: id,
		Name: payload.Name,
		FullyQualifiedName: fmt.Sprintf("%v.%v", payload.DatabaseSchema, payload.Name),
		DisplayName: payload.DisplayName,
		Description: payload.Description,
		StoredProcedureCode: payload.StoredProcedureCode,
		StoredProcedureType: payload.StoredProcedureType,
		ServiceType: dbservice.ServiceType,
		Service: dbserviceEntityRef,
		Database: databaseEntityRef,
		DatabaseSchema: databaseSchemaEntityRef,
		Deleted: false,
	}

	entity := &dataModels.StoredProcedureEntity{
		ID: id,
		Name: payload.Name,
		Json: storedProcedure,
		UpdatedAt: now,
		Deleted: false,
	}

	storedProcedureEntity, err := s.StoredProcedureEntityRepository.InsertStoredProcedureEntity(entity)
	return storedProcedureEntity, err
}

func (s *StoredProcedureEntityService) DeleteStoredProcedureEntityById(id string) error {
	err := s.StoredProcedureEntityRepository.DeleteStoredProcedureEntityById(id)
	return err
}

func (s *StoredProcedureEntityService) DeleteStoredProcedureEntityByFqn(fqn string) error {
	err := s.StoredProcedureEntityRepository.DeleteStoredProcedureEntityByFqn(fqn)
	return err
}
