package constants

const (
	DBInitializationError = "failed to connect with the postgres %s"
	ReadConfigFailedError = "failed to read the config file %s"
	UnmarshalFailedError  = "failed to unmarshal the config file %s"
	MigrationFailedError  = "failed to migrate the database %s"
	DBInstanceFailedError = "failed to initiate the db instance"
)
