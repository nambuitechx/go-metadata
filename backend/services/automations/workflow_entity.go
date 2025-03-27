package services

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	jsonpatch "github.com/evanphx/json-patch"
	baseModels "github.com/nambuitechx/go-metadata/models/base"
	automationsModels "github.com/nambuitechx/go-metadata/models/automations"
	automationsRepositories "github.com/nambuitechx/go-metadata/repositories/automations"
)

type WorkflowEntityService struct {
	WorkflowEntityRepository *automationsRepositories.WorkflowEntityRepository
}

func NewWorkflowEntityService(workflowEntityRepository *automationsRepositories.WorkflowEntityRepository) *WorkflowEntityService {
	return &WorkflowEntityService{ WorkflowEntityRepository: workflowEntityRepository }
}

func (s *WorkflowEntityService) Health() string {
	return "Workflow service is available"
}

func (s *WorkflowEntityService) GetAllWorkflowEntities(limit int, offset int) ([]automationsModels.WorkflowEntity, error) {
	workflowEntity, err := s.WorkflowEntityRepository.SelectWorkflowEntities(limit, offset)
	return workflowEntity, err
}

func (s *WorkflowEntityService) GetCountTableEntities() (*baseModels.EntityTotal, error) {
	entityTotal, err := s.WorkflowEntityRepository.SelectCountWorkflowEntities()
	return entityTotal, err
}

func (s *WorkflowEntityService) GetWorkflowEntityById(id string) (*automationsModels.WorkflowEntity, error) {
	workflowEntity, err := s.WorkflowEntityRepository.SelectWorkflowEntityById(id)
	return workflowEntity, err
}

func (s *WorkflowEntityService) GetWorkflowEntityByFqn(fqn string) (*automationsModels.WorkflowEntity, error) {
	workflowEntity, err := s.WorkflowEntityRepository.SelectWorkflowEntityByFqn(fqn)
	return workflowEntity, err
}

func (s *WorkflowEntityService) CreateWorkflowEntity(payload *automationsModels.CreateWorkflowRequest) (*automationsModels.WorkflowEntity, error) {
	id := uuid.NewString()
	now := time.Now().Unix()

	workflow := &automationsModels.Workflow{
		ID: id,
		Name: payload.Name,
		FullyQualifiedName: payload.Name,
		DisplayName: payload.DisplayName,
		Description: payload.Description,
		WorkflowType: payload.WorkflowType,
		Status: payload.Status,
		Request: payload.Request,
		Response: payload.Response,
		Deleted: false,
	}

	entity := &automationsModels.WorkflowEntity{
		ID: id,
		Name: payload.Name,
		WorkflowType: payload.WorkflowType,
		Status: payload.Status,
		Json: workflow,
		UpdatedAt: now,
		Deleted: false,
	}

	workflowEntity, err := s.WorkflowEntityRepository.InsertWorkflowEntity(entity)
	return workflowEntity, err
}

func (s *WorkflowEntityService) CreateOrUpdateWorkflowEntity(payload *automationsModels.CreateWorkflowRequest) (*automationsModels.WorkflowEntity, error) {
	exist, err := s.WorkflowEntityRepository.SelectWorkflowEntityByFqn(payload.Name)

	if err == nil {
		updated, err := s.WorkflowEntityRepository.UpdateWorkflowEntity(exist)
		return updated, err
	}

	id := uuid.NewString()
	now := time.Now().Unix()

	workflow := &automationsModels.Workflow{
		ID: id,
		Name: payload.Name,
		FullyQualifiedName: payload.Name,
		DisplayName: payload.DisplayName,
		Description: payload.Description,
		WorkflowType: payload.WorkflowType,
		Status: payload.Status,
		Request: payload.Request,
		Response: payload.Response,
		Deleted: false,
	}

	entity := &automationsModels.WorkflowEntity{
		ID: id,
		Name: payload.Name,
		WorkflowType: payload.WorkflowType,
		Status: payload.Status,
		Json: workflow,
		UpdatedAt: now,
		Deleted: false,
	}

	workflowEntity, err := s.WorkflowEntityRepository.InsertWorkflowEntity(entity)
	return workflowEntity, err
}

func (s *WorkflowEntityService) UpdateWorkflowEntity(exist *automationsModels.WorkflowEntity) (*automationsModels.WorkflowEntity, error) {
	entity, err := s.WorkflowEntityRepository.UpdateWorkflowEntity(exist)
	return entity, err
}

func (s *WorkflowEntityService) PatchWorkflowEntity(exist *automationsModels.WorkflowEntity, payload []baseModels.JsonPatchOperation) (*automationsModels.WorkflowEntity, error) {
	// Prepare patch
	jsonPatch, jsonPatchErr := json.Marshal(payload)

	if jsonPatchErr != nil {
		return nil, jsonPatchErr
	}

	patch, patchErr := jsonpatch.DecodePatch(jsonPatch)

	if patchErr != nil {
		return nil, patchErr
	}

	workflowEntityJson, workflowEntityJsonErr := json.Marshal(exist.Json)

	if workflowEntityJsonErr != nil {
		return nil, workflowEntityJsonErr
	}

	// Apply patch
	modified, modifiedErr := patch.Apply(workflowEntityJson)

	if modifiedErr != nil {
		return nil, modifiedErr
	}

	var modifiedWorkflow automationsModels.Workflow

	unmarshalModifiedWorkflowErr := json.Unmarshal(modified, &modifiedWorkflow)

	if unmarshalModifiedWorkflowErr != nil {
		return nil, unmarshalModifiedWorkflowErr
	}

	// Update
	exist.WorkflowType = modifiedWorkflow.WorkflowType
	exist.Status = modifiedWorkflow.Status
	exist.Json = &modifiedWorkflow

	updatedWorkflowEntity, updatedWorkflowEntityErr := s.WorkflowEntityRepository.UpdateWorkflowEntity(exist)

	if updatedWorkflowEntityErr != nil {
		return nil, updatedWorkflowEntityErr
	}

	return updatedWorkflowEntity, nil
}

func (s *WorkflowEntityService) DeleteWorkflowEntityById(id string) error {
	err := s.WorkflowEntityRepository.DeleteWorkflowEntityById(id)
	return err
}

func (s *WorkflowEntityService) DeleteWorkflowEntityByFqn(fqn string) error {
	err := s.WorkflowEntityRepository.DeleteWorkflowEntityByFqn(fqn)
	return err
}


/*
 *
 *	Test connection section
 *
 */ 
// func TestConnection(testServiceConnection *automationsModels.TestServiceConnection) bool {
// 	connection := testServiceConnection.Connection

// 	idx, serviceTypeErr := servicesModels.ValidateServiceType(testServiceConnection.ConnectionType)

// 	if serviceTypeErr != nil {
// 		return false
// 	}

// 	if idx == 0 {
// 		return TestPostgresConnection(connection.Config)
// 	} else if idx == 1 {
// 		return TestMysqlConnection(connection.Config)
// 	} else {
// 		return false
// 	}
// }

// func TestPostgresConnection(conn interface{}) bool {
// 	bytes, err := json.Marshal(conn)

// 	if err != nil {
// 		return false
// 	}

// 	var c *servicesModels.PostgresConnection = &servicesModels.PostgresConnection{}

// 	if err := json.Unmarshal(bytes, c); err != nil {
// 		return false
// 	}

// 	hostPort := strings.Split(c.HostPort, ":")

// 	if len(hostPort) < 2 {
// 		return false
// 	}

// 	host := hostPort[0]
// 	port := hostPort[1]
// 	user := c.Username
// 	password := c.AuthType["password"]
// 	dbName := c.Database

// 	db, err := sql.Open(
// 		"pgx",
// 		fmt.Sprintf(
// 			"postgres://%v:%v@%v:%v/%v",
// 			user,
// 			password,
// 			host,
// 			port,
// 			dbName,
// 		),
// 	)

//     if err != nil {
//         return false
//     }

// 	pingErr := db.Ping()

// 	if pingErr != nil {
//         return false
//     }

// 	closeErr := db.Close()
// 	return closeErr == nil
// }

// func TestMysqlConnection(conn interface{}) bool {
// 	bytes, err := json.Marshal(conn)
	
// 	if err != nil {
// 		return false
// 	}

// 	var c *servicesModels.MysqlConnection = &servicesModels.MysqlConnection{}

// 	if err := json.Unmarshal(bytes, c); err != nil {
// 		return false
// 	}

// 	hostPort := strings.Split(c.HostPort, ":")

// 	if len(hostPort) < 2 {
// 		return false
// 	}

// 	host := hostPort[0]
// 	port := hostPort[1]
// 	user := c.Username
// 	password := c.AuthType["password"]
// 	dbName := c.DatabaseName

// 	db, err := sql.Open(
// 		"mysql",
// 		fmt.Sprintf(
// 			"mysql://%v:%v@%v:%v/%v",
// 			user,
// 			password,
// 			host,
// 			port,
// 			dbName,
// 		),
// 	)

//     if err != nil {
//         return false
//     }

// 	pingErr := db.Ping()

// 	if pingErr != nil {
//         return false
//     } 

// 	closeErr := db.Close()
// 	return closeErr == nil
// }
