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
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS dbservice_entity;
DROP TABLE IF EXISTS database_entity;
DROP TABLE IF EXISTS database_schema_entity;
DROP TABLE IF EXISTS table_entity;
-- +goose StatementEnd
