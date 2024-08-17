package main

import "github.com/dangquanghung/go-ecommerce-backend-api/internal/initialize"

func main() {
	// r := routers.NewServer()
	// r.Run(":8002") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	initialize.Run()
}
