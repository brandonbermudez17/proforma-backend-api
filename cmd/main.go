package main

import (
	"proforma-backend-api/pkg/database"
	"proforma-backend-api/pkg/logger"
	"proforma-backend-api/pkg/server"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func readConfig() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("toml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(err)
	}
}

func main() {
	//reading configuration file
	readConfig()

	//initialize logger
	log := logger.ProvideLogger()
	defer log.Sync()

	//initialize database connection
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Panic("Failed connect to database", zap.Error(err))
	}

	host := viper.GetString("host")
	port := viper.GetUint("port")

	sv := server.New(host, port, db, log)
	sv.Run()

}
