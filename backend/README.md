## Backend server for go-metadata


### APIs


#### Database Services

- [] List database services
GET /v1/services/databaseServices
REQUEST
QUERY-STRING PARAMETERS
    fields (string):        Fields requested in the returned resource. Ex: pipelines,owners,tags,domain
    domain (string):        Filter services by domain. Ex: Marketing
    limit (int32):          Default: 10. Min: 0. Max: 1000000
    before (string):        Returns list of database services before this cursor
    after (string):         Returns list of database services after this cursor
    include (enum):         Default: non-deleted. Allowed: all | deleted | non-deleted

- [] Create database service
POST /v1/services/databaseServices

- [] Create or update database service
PUT /v1/services/databaseServices

- [] Get database service by id
GET /v1/services/databaseServices/{id}
REQUEST
PATH PARAMETERS
    id (string):            Id of the database service
QUERY-STRING PARAMETERS
    fields (string):        Fields requested in the returned resource. Ex: pipelines,owners,tags,domain
    include (enum):         Default: non-deleted. Allowed: all | deleted | non-deleted

- [] Update a database service by id
PATCH /v1/services/databaseServices/{id}

- [] Delete a database service by id
DELETE /v1/services/databaseServices/{id}
REQUEST
PATH PARAMETERS
    id (string):            Id of the database service
QUERY-STRING PARAMETERS
    hardDelete (boolean):   Hard delete the entity. (Default = false)
    recursive (boolean):    Recursively delete this entity and it's children. (Default false)

- [] Get database service by name (For database service, name is also fullyQualifiedName)
GET /v1/services/databaseServices/name/{name}
REQUEST
PATH PARAMETERS
    name (string):          Name of the database service
QUERY-STRING PARAMETERS
    fields (string):        Fields requested in the returned resource. Ex: pipelines,owners,tags,domain
    include (enum):         Default: non-deleted. Allowed: all | deleted | non-deleted

- [] Update a database service by name (For database service, name is also fullyQualifiedName)
PATCH /v1/services/databaseServices/name/{name}

- [] Delete a database service by name (For database service, name is also fullyQualifiedName)
DELETE /v1/services/databaseServices/name/{name}
REQUEST
PATH PARAMETERS
    name (string):          Name of the database service
QUERY-STRING PARAMETERS
    hardDelete (boolean):   Hard delete the entity. (Default = false)
    recursive (boolean):    Recursively delete this entity and it's children. (Default false)

- [] Restore a soft deleted database service
PUT /v1/services/databaseServices/restore
REQUEST
REQUEST BODY
    { "id": "..." }


#### Databases

- [] List databases
GET /v1/databases
REQUEST
QUERY-STRING PARAMETERS
    fields (string):        Fields requested in the returned resource. Ex: owners,databaseSchemas,usageSummary,location,tags,extension,domain,sourceHash
    service (string):       Filter databases by service name. Ex: snowflakeWestCoast
    limit (int32):          Default: 10. Min: 0. Max: 1000000
    before (string):        Returns list of databases before this cursor
    after (string):         Returns list of databases after this cursor
    include (enum):         Default: non-deleted. Allowed: all | deleted | non-deleted

- [] Create database
POST /v1/databases

- [] Create or update database
PUT /v1/databases

- [] Get database by id
GET /v1/databases/{id}
REQUEST
PATH PARAMETERS
    id (string):            Id of the database
QUERY-STRING PARAMETERS
    fields (string):        Fields requested in the returned resource. Ex: owners,databaseSchemas,usageSummary,location,tags,extension,domain,sourceHash
    include (enum):         Default: non-deleted. Allowed: all | deleted | non-deleted

- [] Update a database by id
PATCH /v1/databases/{id}

- [] Delete a database by id
DELETE /v1/databases/{id}
REQUEST
PATH PARAMETERS
    id (string):            Id of the database
QUERY-STRING PARAMETERS
    hardDelete (boolean):   Hard delete the entity. (Default = false)
    recursive (boolean):    Recursively delete this entity and it's children. (Default false)

- [] Get database by fqn
GET /v1/databases/name/{fqn}
REQUEST
PATH PARAMETERS
    fqn (string):           Fully qualified name of the database
QUERY-STRING PARAMETERS
    fields (string):        Fields requested in the returned resource. Ex: owners,databaseSchemas,usageSummary,location,tags,extension,domain,sourceHash
    include (enum):         Default: non-deleted. Allowed: all | deleted | non-deleted

- [] Update a database by fqn
PATCH /v1/databases/name/{fqn}

- [] Delete a database by fqn
DELETE /v1/databases/name/{fqn}
REQUEST
PATH PARAMETERS
    fqn (string):           Fully qualified name of the database
QUERY-STRING PARAMETERS
    hardDelete (boolean):   Hard delete the entity. (Default = false)
    recursive (boolean):    Recursively delete this entity and it's children. (Default false)

- [] Restore a soft deleted database
PUT /v1/databases/restore
REQUEST
REQUEST BODY
    { "id": "..." }


#### Database Schemas

- [] List database schemas
GET /v1/databaseSchemas
REQUEST
QUERY-STRING PARAMETERS
    fields (string):        Fields requested in the returned resource. Ex: owners,tables,usageSummary,tags,extension,domain,sourceHash
    database (string):      Filter schemas by database name. Ex: customerDatabase
    limit (int32):          Default: 10. Min: 0. Max: 1000000
    before (string):        Returns list of databases before this cursor
    after (string):         Returns list of databases after this cursor
    include (enum):         Default: non-deleted. Allowed: all | deleted | non-deleted

- [] Create database schema
POST /v1/databaseSchemas

- [] Create or update database schema
PUT /v1/databaseSchemas

- [] Get database schema by id
GET /v1/databaseSchemas/{id}
REQUEST
PATH PARAMETERS
    id (string):            Id of the database schema
QUERY-STRING PARAMETERS
    fields (string):        Fields requested in the returned resource. Ex: owners,tables,usageSummary,tags,extension,domain,sourceHash
    include (enum):         Default: non-deleted. Allowed: all | deleted | non-deleted

- [] Update a database schema by id
PATCH /v1/databaseSchemas/{id}

- [] Delete a database schema by id
DELETE /v1/databaseSchemas/{id}
REQUEST
PATH PARAMETERS
    id (string):            Id of the database schema
QUERY-STRING PARAMETERS
    hardDelete (boolean):   Hard delete the entity. (Default = false)
    recursive (boolean):    Recursively delete this entity and it's children. (Default false)

- [] Get database schema by fqn
GET /v1/databaseSchemas/name/{fqn}
REQUEST
PATH PARAMETERS
    fqn (string):           Fully qualified name of the database schema
QUERY-STRING PARAMETERS
    fields (string):        Fields requested in the returned resource. Ex: owners,tables,usageSummary,tags,extension,domain,sourceHash
    include (enum):         Default: non-deleted. Allowed: all | deleted | non-deleted

- [] Update a database schema by fqn
PATCH /v1/databaseSchemas/name/{fqn}

- [] Delete a database schema by fqn
DELETE /v1/databaseSchemas/name/{fqn}
REQUEST
PATH PARAMETERS
    fqn (string):           Fully qualified name of the database schema
QUERY-STRING PARAMETERS
    hardDelete (boolean):   Hard delete the entity. (Default = false)
    recursive (boolean):    Recursively delete this entity and it's children. (Default false)

- [] Restore a soft deleted database schema
PUT /v1/databaseSchemas/restore
REQUEST
REQUEST BODY
    { "id": "..." }


#### Tables

- [] List tables
GET /v1/tables
REQUEST
QUERY-STRING PARAMETERS
    fields (string):        Fields requested in the returned resource. Ex: owners,tables,usageSummary,tags,extension,domain,sourceHash
    database (string):      Filter schemas by database fully qualified name. Ex: snowflakeWestCoast.financeDB
    databaseSchema (str):   Filter tables by databaseSchema fully qualified name. Ex: snowflakeWestCoast.financeDB.schema
    includeEmptyTestSuite:  Include tables with an empty test suite. Default: true
    limit (int32):          Default: 10. Min: 0. Max: 1000000
    before (string):        Returns list of databases before this cursor
    after (string):         Returns list of databases after this cursor
    include (enum):         Default: non-deleted. Allowed: all | deleted | non-deleted

- [] Create table
POST /v1/tables

- [] Create or update table
PUT /v1/tables

- [] Get table by id
GET /v1/tables/{id}
REQUEST
PATH PARAMETERS
    id (string):            Id of the table
QUERY-STRING PARAMETERS
    fields (string):        Fields requested in the returned resource. Ex: tableConstraints,tablePartition,usageSummary,owners,customMetrics,columns,tags,followers,joins,schemaDefinition,dataModel,extension,testSuite,domain,dataProducts,lifeCycle,sourceHash
    include (enum):         Default: non-deleted. Allowed: all | deleted | non-deleted

- [] Update a table by id
PATCH /v1/tables/{id}

- [] Delete a table by id
DELETE /v1/tables/{id}
REQUEST
PATH PARAMETERS
    id (string):            Id of the table
QUERY-STRING PARAMETERS
    hardDelete (boolean):   Hard delete the entity. (Default = false)
    recursive (boolean):    Recursively delete this entity and it's children. (Default false)

- [] Get table by fqn
GET /v1/tables/name/{fqn}
REQUEST
PATH PARAMETERS
    fqn (string):           Fully qualified name of the table
QUERY-STRING PARAMETERS
    fields (string):        Fields requested in the returned resource. Ex: tableConstraints,tablePartition,usageSummary,owners,customMetrics,columns,tags,followers,joins,schemaDefinition,dataModel,extension,testSuite,domain,dataProducts,lifeCycle,sourceHash
    include (enum):         Default: non-deleted. Allowed: all | deleted | non-deleted

- [] Update a table by fqn
PATCH /v1/tables/name/{fqn}

- [] Delete a table by fqn
DELETE /v1/tables/name/{fqn}
REQUEST
PATH PARAMETERS
    fqn (string):           Fully qualified name of the table
QUERY-STRING PARAMETERS
    hardDelete (boolean):   Hard delete the entity. (Default = false)
    recursive (boolean):    Recursively delete this entity and it's children. (Default false)

- [] Restore a soft deleted table
PUT /v1/tables/restore
REQUEST
REQUEST BODY
    { "id": "..." }
