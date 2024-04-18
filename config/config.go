package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var c config

type AppConfig struct {
	Name     string         `mapstructure:"name"`
	Address  string         `mapstructure:"address"`
	Port     string         `mapstructure:"port"`
	Transfer TransferConfig `mapstructure:"transfer"`
}

type TransferConfig struct {
	Registered   AccountConfig `mapstructure:"registered"`
	Unregistered AccountConfig `mapstructure:"unregistered"`
}

type AccountConfig struct {
	BalanceLimit             int32   `mapstructure:"balance_limit"`
	CreditCountDailyLimit    int16   `mapstructure:"credit_count_daily_limit"`
	CreditCountMonthlyLimit  int16   `mapstructure:"credit_count_monthly_limit"`
	CreditAmountDailyLimit   float64 `mapstructure:"credit_amount_daily_limit"`
	CreditAmountMonthlyLimit float64 `mapstructure:"credit_amount_monthly_limit"`
}

type StatsDConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type DatadogConfig struct {
	Enabled bool   `mapstructure:"enabled"`
	Env     string `mapstructure:"env"`
	Name    string `mapstructure:"service"`
	Version string `mapstructure:"version"`
}

// type SentryConfig struct {
// 	Enabled bool   `mapstructure:"enabled"`
// 	DSN     string `mapstructure:"dsn"`
// }

type config struct {
	App      AppConfig     `mapstructure:"app"`
	StatsD   StatsDConfig  `mapstructure:"statsd"`
	Datadog  DatadogConfig `mapstructure:"datadog"`
	DBConfig DBConfig      `mapstructure:"database"`
}

func Load(cfgName, path string) error {
	viper.SetConfigName(cfgName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	_ = viper.ReadInConfig()
	return viper.Unmarshal(&c)
}

func GetAppConfig() AppConfig {
	return c.App
}

func GetDatadogConfig() DatadogConfig {
	return c.Datadog
}

// func GetSentry() SentryConfig {
// 	return c.Sentry
// }

func GetDBConfig() DBConfig {
	return c.DBConfig
}

func GetStatsDAddress() string {
	return fmt.Sprintf("%s:%s", c.StatsD.Host, c.StatsD.Port)
}
