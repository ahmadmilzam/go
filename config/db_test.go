package config_test

import (
	"testing"

	"github.com/ahmadmilzam/go/config"
	"github.com/stretchr/testify/assert"
)

func TestGetConnectionURI(t *testing.T) {
	cfg := config.DBConfig{
		Host:     "localhost",
		Name:     "ewallet",
		Username: "postgres",
		Password: "password",
		Port:     "5432",
	}

	assert.Equal(t, "postgresql://postgres:password@localhost:5432/ewallet?sslmode=disable", cfg.GetConnectionURI())
}
