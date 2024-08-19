package initialize

import (
	"fmt"

	"github.com/dangquanghung/go-ecommerce-backend-api/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
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
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
}
