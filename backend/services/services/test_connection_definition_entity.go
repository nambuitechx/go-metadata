package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	baseModels "github.com/nambuitechx/go-metadata/models/base"
	servicesModels "github.com/nambuitechx/go-metadata/models/services"
	servicesRepositories "github.com/nambuitechx/go-metadata/repositories/services"
)

type TestConnectionDefinitionEntityService struct {
	TestConnectionDefinitionEntityRepository *servicesRepositories.TestConnectionDefinitionEntityRepository
}

func NewTestConnectionDefinitionEntityService(testConnectionDefinitionEntityRepository *servicesRepositories.TestConnectionDefinitionEntityRepository) *TestConnectionDefinitionEntityService {
	service := &TestConnectionDefinitionEntityService{ TestConnectionDefinitionEntityRepository: testConnectionDefinitionEntityRepository }
	service.InitTestConnectionDefinitions()
	return service
}

func (s *TestConnectionDefinitionEntityService) Health() string {
	return "Test connection definition service is available"
}

func (s *TestConnectionDefinitionEntityService) GetAllTestConnectionDefinitionEntities(limit int, offset int) ([]servicesModels.TestConnectionDefinitionEntity, error) {
	testConnectionDefinitionEntity, err := s.TestConnectionDefinitionEntityRepository.SelectTestConnectionDefinitionEntities(limit, offset)
	return testConnectionDefinitionEntity, err
}

func (s *TestConnectionDefinitionEntityService) GetCountTestConnectionDefinitionEntities() (*baseModels.EntityTotal, error) {
	entityTotal, err := s.TestConnectionDefinitionEntityRepository.SelectCountTestConnectionDefinitionEntities()
	return entityTotal, err
}

func (s *TestConnectionDefinitionEntityService) GetTestConnectionDefinitionEntityById(id string) (*servicesModels.TestConnectionDefinitionEntity, error) {
	testConnectionDefinitionEntity, err := s.TestConnectionDefinitionEntityRepository.SelectTestConnectionDefinitionEntityById(id)
	return testConnectionDefinitionEntity, err
}

func (s *TestConnectionDefinitionEntityService) GetTestConnectionDefinitionEntityByFqn(fqn string) (*servicesModels.TestConnectionDefinitionEntity, error) {
	testConnectionDefinitionEntity, err := s.TestConnectionDefinitionEntityRepository.SelectTestConnectionDefinitionEntityByFqn(fqn)
	return testConnectionDefinitionEntity, err
}

func (s *TestConnectionDefinitionEntityService) InitTestConnectionDefinitions() {
	path := "./json/data/test-connections/database"
	files, err := os.ReadDir(path)

	if err != nil {
		log.Fatalln(err)
	}

    for _, entry := range files {
		jsonFile, openFileErr := os.Open(fmt.Sprintf("%v/%v", path, entry.Name()))

		if openFileErr != nil {
			log.Fatalln(openFileErr)
		}

		defer jsonFile.Close()

		byteValue, readErr := io.ReadAll(jsonFile)

		if readErr != nil {
			log.Fatalln(readErr)
		}

		var data TestConnectionDefinitionData

		if err := json.Unmarshal(byteValue, &data); err != nil {
			log.Fatalln(err)
		}

		id := uuid.NewString()
		fullyQualifiedName := fmt.Sprintf("%v.%v", data.Name, servicesModels.TestConnectionDefinitionString)
		now := time.Now().Unix()

		_, err := s.TestConnectionDefinitionEntityRepository.SelectTestConnectionDefinitionEntityByFqn(fullyQualifiedName)

		if err != nil {
			log.Printf("========== Init test connection for %v", data.Name)

			testConnetionDefinition := &servicesModels.TestConnectionDefinition {
				ID: id,
				Name: data.Name,
				FullyQualifiedName: fullyQualifiedName,
				DisplayName: data.DisplayName,
				Description: data.Description,
				Steps: data.Steps,
				Deleted: false,
			}

			entity := &servicesModels.TestConnectionDefinitionEntity {
				ID: id,
				Name: data.Name,
				FullyQualifiedName: fullyQualifiedName,
				Json: testConnetionDefinition,
				UpdatedAt: now,
				Deleted: false,
			}

			if _, err := s.TestConnectionDefinitionEntityRepository.InsertTestConnectionDefinitionEntity(entity); err != nil {
				log.Fatalln(err)
			}
		}
    }
}

type TestConnectionDefinitionData struct {
	Name			string									`json:"name"`
	DisplayName		string									`json:"displayName"`
	Description		string									`json:"description"`
	Steps			[]*servicesModels.TestConnectionStep	`json:"steps"`
}
