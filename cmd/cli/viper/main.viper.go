package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration %w ", err))
	}

	// read server configuration

	// fmt.Println("Server Port:: ", viper.GetInt("server.port"))
	// fmt.Println("Server Port:: ", viper.GetString("security.jwt.key"))

	// configuration structure
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}

	fmt.Println("Config Port:: ", config.Server.Port)

	for _, db := range config.Databases {
		fmt.Printf("databases user: %s, password: %s, host: %s \n", db.User, db.Password, db.Host)
	}
}
