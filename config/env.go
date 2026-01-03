package config

import "os"

var Env envManager

type envManager struct {
	APP_PORT   string
	APP_ENV    string
	APP_LOCALE string
	APP_SECRET string
	APP_URL    string

	APP_DB_HOST     string
	APP_DB_PORT     string
	APP_DB_NAME     string
	APP_DB_USER     string
	APP_DB_PASS     string
	APP_DB_SCHEMA   string
	APP_DB_SSLMODE  string
	APP_DB_TimeZone string
}

func (m *envManager) LoadEnv() {
	m.APP_PORT = m.Get("APP_PORT", "4000")
	m.APP_ENV = m.Get("APP_ENV", "local")
	m.APP_LOCALE = m.Get("APP_LOCALE", "en")
	m.APP_SECRET = m.Get("APP_SECRET", "secret")
	m.APP_URL = m.Get("APP_URL", "localhost:4000")

	m.APP_DB_HOST = m.Get("APP_DB_HOST", "localhost")
	m.APP_DB_PORT = m.Get("APP_DB_PORT", "5432")
	m.APP_DB_NAME = m.Get("APP_DB_NAME", "go-fiber-react")
	m.APP_DB_USER = m.Get("APP_DB_USER", "postgres")
	m.APP_DB_PASS = m.Get("APP_DB_PASS", "\"\"")
	m.APP_DB_SCHEMA = m.Get("APP_DB_SCHEMA", "public")
	m.APP_DB_SSLMODE = m.Get("APP_DB_SSLMODE", "disable")
	m.APP_DB_TimeZone = m.Get("APP_DB_TimeZone", "Asia/Jakarta")
}

func (*envManager) Get(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}
