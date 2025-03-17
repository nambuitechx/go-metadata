-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS dbservice_entity(
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(256) UNIQUE NOT NULL,
    servicetype VARCHAR(256) NOT NULL,
    json JSONB NOT NULL,
    updatedat BIGINT NOT NULL,
    updatedby VARCHAR(256),
    deleted BOOLEAN NOT NULL,
    namehash VARCHAR(256)
);
CREATE TABLE IF NOT EXISTS database_entity(
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(256) UNIQUE NOT NULL,
    json JSONB NOT NULL,
    updatedat BIGINT NOT NULL,
    updatedby VARCHAR(256),
    deleted BOOLEAN NOT NULL,
    fqnhash VARCHAR(256)
);
CREATE TABLE IF NOT EXISTS database_schema_entity(
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(256) UNIQUE NOT NULL,
    json JSONB NOT NULL,
    updatedat BIGINT NOT NULL,
    updatedby VARCHAR(256),
    deleted BOOLEAN NOT NULL,
    fqnhash VARCHAR(256)
);
CREATE TABLE IF NOT EXISTS table_entity(
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(256) UNIQUE NOT NULL,
    json JSONB NOT NULL,
    updatedat BIGINT NOT NULL,
    updatedby VARCHAR(256),
    deleted BOOLEAN NOT NULL,
    fqnhash VARCHAR(256)
);
CREATE TABLE IF NOT EXISTS automations_workflow(
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(256) UNIQUE NOT NULL,
    workflowtype VARCHAR(256) NOT NULL,
    status VARCHAR(256) NOT NULL,
    json JSONB NOT NULL,
    updatedat BIGINT NOT NULL,
    updatedby VARCHAR(256),
    deleted BOOLEAN NOT NULL,
    namehash VARCHAR(256)
);
CREATE TABLE IF NOT EXISTS bot_entity(
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(256) UNIQUE NOT NULL,
    json JSONB NOT NULL,
    updatedat BIGINT NOT NULL,
    updatedby VARCHAR(256),
    deleted BOOLEAN NOT NULL,
    namehash VARCHAR(256)
);
CREATE TABLE IF NOT EXISTS change_event(
    eventtype VARCHAR(36) NOT NULL,
    entitytype VARCHAR(36) NOT NULL,
    username VARCHAR(256) NOT NULL,
    eventtime BIGINT NOT NULL,
    json JSONB NOT NULL,
    offset INTEGER PRIMARY KEY
);
CREATE TABLE IF NOT EXISTS entity_relationship(
    fromid VARCHAR(36) PRIMARY KEY,
    toid VARCHAR(36) PRIMARY KEY,
    fromentity VARCHAR(256) NOT NULL,
    toentity VARCHAR(256) NOT NULL,
    relation SMALLINT NOT NULL,
    jsonschema VARCHAR(256),
    json JSONB,
    deleted BOOLEAN NOT NULL
);
CREATE TABLE IF NOT EXISTS ingestion_pipeline_entity(
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(256) UNIQUE NOT NULL,
    json JSONB NOT NULL,
    updatedat BIGINT NOT NULL,
    updatedby VARCHAR(256),
    deleted BOOLEAN NOT NULL,
    fqnhash VARCHAR(256),
    timestamp BIGINT,
    apptype VARCHAR(256),
    pipelinetype VARCHAR(256)
);
CREATE TABLE IF NOT EXISTS pipeline_entity(
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(256) UNIQUE NOT NULL,
    json JSONB NOT NULL,
    updatedat BIGINT NOT NULL,
    updatedby VARCHAR(256),
    deleted BOOLEAN NOT NULL,
    fqnhash VARCHAR(256)
);
CREATE TABLE IF NOT EXISTS pipeline_service_entity(
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(256) UNIQUE NOT NULL,
    servicetype VARCHAR(256) NOT NULL,
    json JSONB NOT NULL,
    updatedat BIGINT NOT NULL,
    updatedby VARCHAR(256),
    deleted BOOLEAN NOT NULL,
    namehash VARCHAR(256)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS dbservice_entity;
DROP TABLE IF EXISTS database_entity;
DROP TABLE IF EXISTS database_schema_entity;
DROP TABLE IF EXISTS table_entity;
DROP TABLE IF EXISTS automations_workflow;
DROP TABLE IF EXISTS bot_entity;
DROP TABLE IF EXISTS change_event;
DROP TABLE IF EXISTS entity_relationship;
DROP TABLE IF EXISTS ingestion_pipeline_entity;
DROP TABLE IF EXISTS pipeline_entity;
DROP TABLE IF EXISTS pipeline_service_entity;
-- +goose StatementEnd
