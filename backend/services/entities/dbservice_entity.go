package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/nambuitechx/go-metadata/models/entities"
	"github.com/nambuitechx/go-metadata/repositories/entities"
)

type DBServiceEntityService struct {
	DBServiceEnityRepository *repositories.DBServiceEntityRepository
}

func NewDBServiceEntityService(dbserviceEntityRepository *repositories.DBServiceEntityRepository) *DBServiceEntityService {
	return &DBServiceEntityService{ DBServiceEnityRepository: dbserviceEntityRepository }
}

func (s *DBServiceEntityService) Health() string {
	return "DBService service is available"
}

func (s *DBServiceEntityService) GetAllDBServiceEntities(limit int, offset int) ([]models.DBServiceEntity, error) {
	dbserviceEntity, err := s.DBServiceEnityRepository.SelectDBServiceEntities(limit, offset)
	return dbserviceEntity, err
}

func (s *DBServiceEntityService) GetDBServiceEntityById(id string) (*models.DBServiceEntity, error) {
	dbserviceEntity, err := s.DBServiceEnityRepository.SelectDBServiceEntityById(id)
	return dbserviceEntity, err
}

func (s *DBServiceEntityService) GetDBServiceEntityByFqn(fqn string) (*models.DBServiceEntity, error) {
	dbserviceEntity, err := s.DBServiceEnityRepository.SelectDBServiceEntityByFqn(fqn)
	return dbserviceEntity, err
}

func (s *DBServiceEntityService) CreateDBServiceEntity(payload *models.CreateDBServiceEntityPayload) (*models.DBServiceEntity, error) {
	id := uuid.NewString()
	now := time.Now().Unix()

	dbservice := &models.DBService{
		ID: id,
		Name: payload.Name,
		FullyQualifiedName: payload.Name,
		DisplayName: payload.DisplayName,
		Description: payload.Description,
		ServiceType: payload.ServiceType,
		Connection: payload.Connection,
		Deleted: false,
	}

	entity := &models.DBServiceEntity{
		ID: id,
		Name: payload.Name,
		ServiceType: payload.ServiceType,
		Json: dbservice,
		UpdatedAt: now,
		Deleted: false,
	}

	dbserviceEntity, err := s.DBServiceEnityRepository.InsertDBServiceEntity(entity)
	return dbserviceEntity, err
}

func (s *DBServiceEntityService) DeleteDBServiceEntityById(id string) error {
	err := s.DBServiceEnityRepository.DeleteDBServiceEntityById(id)
	return err
}

func (s *DBServiceEntityService) DeleteDBServiceEntityByFqn(fqn string) error {
	err := s.DBServiceEnityRepository.DeleteDBServiceEntityByFqn(fqn)
	return err
}
