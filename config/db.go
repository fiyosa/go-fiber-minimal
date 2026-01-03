package config

var Db dbManager

type dbManager struct {
	Host     string
	User     string
	Password string
	Name     string
	Schema   string
	Port     string
	SSLMode  string
	TimeZone string

	APP_ENV string
}

func (m *dbManager) Init() {
	m.Host = Env.APP_DB_HOST
	m.User = Env.APP_DB_USER
	m.Password = Env.APP_DB_PASS
	m.Name = Env.APP_DB_NAME
	m.Schema = Env.APP_DB_SCHEMA
	m.Port = Env.APP_DB_PORT
	m.SSLMode = Env.APP_DB_SSLMODE
	m.TimeZone = Env.APP_DB_TimeZone

	m.APP_ENV = Env.APP_ENV
}
