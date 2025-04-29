package constants

const (
	DSNString = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"
)

const (
	BasePath = "../../config"
	RootPath = "./src/config"
)

const (
	PostgresConfigName = "postgres"
	ConfigType         = "yaml"
)

const (
	MigrationSuccessMessage = "database is migrated successfully"
)

const (
	PanCardRegex = "^[A-Z]{0,5}[0-9]{0,4}[A-z]{1}"
)

const (
	JsonBindingFieldError = "unexpected value for the field"
)

// Database table name & field names for users
const (
	Username          = "username"
	Name              = "name"
	Email             = "email"
	PhoneNumber       = "phoneNumber"
	PanCard           = "panCard"
	Password          = "password"
)