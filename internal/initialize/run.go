package initialize

import (
	"fmt"

	"github.com/dangquanghung/go-ecommerce-backend-api/global"
)

func Run() {
	// load configuration
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql --->", m.Username, m.Password)
	InitLogger()
	InitMysql()
	InitRedis()

	r := InitRouter()

	r.Run(":8002")
}
