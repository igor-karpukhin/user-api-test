package config

import "flag"

type ApplicationConfiguration struct {
	HttpHost       string
	HttpPort       uint
	PostgresConfig *PostgresConfiguration
}

func NewApplicationConfiguration() *ApplicationConfiguration {
	return &ApplicationConfiguration{
		HttpHost:       "",
		HttpPort:       0,
		PostgresConfig: &PostgresConfiguration{},
	}
}

func FromFlags() *ApplicationConfiguration {
	c := NewApplicationConfiguration()

	flag.StringVar(&c.HttpHost, "addr", "127.0.0.1", "server ip")
	flag.UintVar(&c.HttpPort, "port", 8080, "server port")
	flag.StringVar(&c.PostgresConfig.Host, "pg.host", "127.0.0.1", "postgres host")
	flag.UintVar(&c.PostgresConfig.Port, "pg.port", 5432, "postgres port")
	flag.StringVar(&c.PostgresConfig.Username, "pg.user", "", "pg user")
	flag.StringVar(&c.PostgresConfig.Password, "pg.pass", "", "pg password")
	flag.StringVar(&c.PostgresConfig.DBName, "pg.dbname", "", "pg db name")
	flag.Parse()

	return c
}
