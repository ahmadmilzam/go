package config

import "fmt"

type DBConfig struct {
	Host       string           `mapstructure:"host"`
	Name       string           `mapstructure:"name"`
	Username   string           `mapstructure:"username"`
	Password   string           `mapstructure:"password"`
	Port       string           `mapstructure:"port"`
	Connection ConnectionConfig `mapstructure:"connection"`
}

type ConnectionConfig struct {
	MaxIdleConn     int `mapstructure:"maxIdleConn"`
	MaxOpenConn     int `mapstructure:"maxOpenConn"`
	MaxLifeTimeConn int `mapstructure:"maxLifeTimeConn"`
}

func (c *DBConfig) GetConnectionURI() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", c.Username, c.Password, c.Host, c.Port, c.Name)
}
