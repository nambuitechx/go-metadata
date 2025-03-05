package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/nambuitechx/go-metadata/models"
	"github.com/nambuitechx/go-metadata/repositories"
)

type DBServiceService struct {
	DBServiceEnityRepository *repositories.DBServiceEntityRepository
}

func NewDBServiceService(dbserviceEnityRepository *repositories.DBServiceEntityRepository) *DBServiceService {
	return &DBServiceService{ DBServiceEnityRepository: dbserviceEnityRepository }
}

func (s *DBServiceService) Health() string {
	return "DBService service is available"
}

func (s *DBServiceService) GetAllDBServiceEntities(limit int, offset int) ([]models.DBServiceEntity, error) {
	dbserviceEntity, err := s.DBServiceEnityRepository.SelectDBServiceEntities(limit, offset)
	return dbserviceEntity, err
}

func (s *DBServiceService) GetDBServiceEntityById(id string) (*models.DBServiceEntity, error) {
	dbserviceEntity, err := s.DBServiceEnityRepository.SelectDBServiceEntityById(id)
	return dbserviceEntity, err
}

func (s *DBServiceService) CreateDBServiceEntity(payload *models.CreateDBServiceEntity) (*models.DBServiceEntity, error) {
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

func (s *DBServiceService) DeleteDBServiceEntityById(id string) error {
	err := s.DBServiceEnityRepository.DeleteDBServiceEntityById(id)
	return err
}
