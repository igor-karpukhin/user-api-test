package config

type PostgresConfiguration struct {
	Host     string
	Port     uint
	Username string
	Password string
	DBName   string
}
