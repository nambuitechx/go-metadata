from metadata.generated.schema.entity.services.databaseService import DatabaseServiceType

INTERNAL_SCHEMA = {
    DatabaseServiceType.Oracle: [
        "ANONYMOUS",  # HTTP access to XDB
        "APPQOSSYS",  # QOS system user
        "AUDSYS",  # audit super user
        "BI",  # Business Intelligence
        "CTXSYS",  # Text
        "DBSNMP",  # SNMP agent for OEM
        "DIP",  # Directory Integration Platform
        "DMSYS",  # Data Mining
        "DVF",  # Database Vault
        "DVSYS",  # Database Vault
        "EXDSYS",  # External ODCI System User
        "EXFSYS",  # Expression Filter
        "GSMADMIN_INTERNAL",  # Global Service Manager
        "GSMCATUSER",  # Global Service Manager
        "GSMUSER",  # Global Service Manager
        "LBACSYS",  # Label Security
        "MDSYS",  # Spatial
        "SPATIAL_CSW_ADMIN",  # Spatial Catalog Services for Web
        "SPATIAL_CSW_ADMIN_USR",  # Spatial
        "SPATIAL_WFS_ADMIN",  # Spatial Web Feature Service
        "SPATIAL_WFS_ADMIN_USR",  # Spatial
        "MGMT_VIEW",  # OEM Database Control
        "MTSSYS",  # MS Transaction Server
        "ODM",  # Data Mining
        "ODM_MTR",  # Data Mining Repository
        "OJVMSYS",  # Java Policy SRO Schema
        "OLAPSYS",  # OLAP catalogs
        "ORACLE_OCM",  # Oracle Configuration Manager User
        "ORDDATA",  # Intermedia
        "ORDPLUGINS",  # Intermedia
        "ORDSYS",  # Intermedia
        "OUTLN",  # Outlines (Plan Stability)
        "SI_INFORMTN_SCHEMA",  # SQL/MM Still Image
        "SYS",
        "SYSBACKUP",
        "SYSDG",
        "SYSKM",
        "SYSMAN",  # Adminstrator OEM
        "SYSTEM",
        "TSMSYS",  # Transparent Session Migration
        "WKPROXY",  # Ultrasearch
        "WKSYS",  # Ultrasearch
        "WMSYS",  # Workspace Manager
        "XDB",  # XML DB
        "XTISYS",  # Time Index
        "AURORA$JIS$UTILITY$",  # JSERV
        "AURORA$ORB$UNAUTHENTICATED",  # JSERV
        "DSSYS",  # Dynamic Services Secured Web Service
        "OSE$HTTP$ADMIN",  # JSERV
        "PERFSTAT",  # STATSPACK
        "TRACESVR",  # Trace server for OEM
        "dbsfwuser",  # Database Firewall
        "dgpdb_int",  # Database Gateway
        "ggsys",  # GoldenGate
        "mddata",  # OLAP Analytic Workspace
        "remote_scheduler_agent",  # Remote Scheduler Agent
        "xs.*null",  # Oracle Text
    ],
    DatabaseServiceType.Postgres: ["information_schema"],
    DatabaseServiceType.Mysql: [
        "information_schema",
        "mysql",
        "performance_schema",
        "sys",
    ],
    DatabaseServiceType.Mssql: [
        "sys",
        "INFORMATION_SCHEMA",
        "guest",
        "db_.*",
    ],
}