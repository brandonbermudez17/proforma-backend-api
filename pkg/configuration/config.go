package configuration

import (
	"fmt"

	"github.com/spf13/viper"
)

type DbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func getConfig() *DbConfig {

	config := &DbConfig{}

	config.Host = viper.GetString("database.server")
	config.Port = viper.GetInt("database.port")
	config.User = viper.GetString("database.user")
	config.Password = viper.GetString("database.password")
	config.Database = viper.GetString("database.name")

	return config
}

func GetConnectionString() string {
	config := getConfig()
	dsn := fmt.Sprintf(
		// "host=%s user=%s password=%s dbname=%s --set=sslmode=verify-full",
		"postgresql://%s:%s@%s:25060/%s?sslmode=require",
		config.User,
		config.Password,
		config.Host,
		config.Database,
	)
	return dsn
}
