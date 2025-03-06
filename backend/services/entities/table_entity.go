package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/nambuitechx/go-metadata/models/entities"
	"github.com/nambuitechx/go-metadata/repositories/entities"
)

type TableEntityService struct {
	DBServiceEntityRepository *repositories.DBServiceEntityRepository
	DatabaseEntityRepository *repositories.DatabaseEntityRepository
	DatabaseSchemaEntityRepository *repositories.DatabaseSchemaEntityRepository
	TableEntityRepository *repositories.TableEntityRepository
}

func NewTableEntityService(
	dbserviceEntityRepository *repositories.DBServiceEntityRepository,
	databaseEntityRepository *repositories.DatabaseEntityRepository,
	databaseSchemaEntityRepository *repositories.DatabaseSchemaEntityRepository,
	tableEntityRepository *repositories.TableEntityRepository,
) *TableEntityService {
	return &TableEntityService{
		DBServiceEntityRepository: dbserviceEntityRepository,
		DatabaseEntityRepository: databaseEntityRepository,
		DatabaseSchemaEntityRepository: databaseSchemaEntityRepository,
		TableEntityRepository: tableEntityRepository,
	}
}

func (s *TableEntityService) Health() string {
	return "Table service is available"
}

func (s *TableEntityService) GetAllTableEntities(limit int, offset int) ([]models.TableEntity, error) {
	tableEntity, err := s.TableEntityRepository.SelectTableEntities(limit, offset)
	return tableEntity, err
}

func (s *TableEntityService) GetTableEntityById(id string) (*models.TableEntity, error) {
	tableEntity, err := s.TableEntityRepository.SelectTableEntityById(id)
	return tableEntity, err
}

func (s *TableEntityService) GetTableEntityByFqn(fqn string) (*models.TableEntity, error) {
	tableEntity, err := s.TableEntityRepository.SelectTableEntityByFqn(fqn)
	return tableEntity, err
}

func (s *TableEntityService) CreateTableEntity(payload *models.CreateTableEntityPayload) (*models.TableEntity, error) {
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

	// Populate table
	table := &models.Table{
		ID: id,
		Name: payload.Name,
		FullyQualifiedName: fmt.Sprintf("%v.%v", payload.DatabaseSchema, payload.Name),
		DisplayName: payload.DisplayName,
		Description: payload.Description,
		ServiceType: dbservice.ServiceType,
		Service: dbserviceEntityRef,
		Database: databaseEntityRef,
		DatabaseSchema: databaseSchemaEntityRef,
		TableType: payload.TableType,
		TableConstraints: payload.TableConstraints,
		Columns: payload.Columns,
		Deleted: false,
	}

	entity := &models.TableEntity{
		ID: id,
		Name: payload.Name,
		Json: table,
		UpdatedAt: now,
		Deleted: false,
	}

	tableEntity, err := s.TableEntityRepository.InsertTableEntity(entity)
	return tableEntity, err
}

func (s *TableEntityService) DeleteTableEntityById(id string) error {
	err := s.TableEntityRepository.DeleteTableEntityById(id)
	return err
}

func (s *TableEntityService) DeleteTableEntityByFqn(fqn string) error {
	err := s.TableEntityRepository.DeleteTableEntityByFqn(fqn)
	return err
}
