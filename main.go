package main

import (
	"Go-Concurrency/config"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func OSFunctions(cmd string) (string, error) {

	// fmt.Println("hello os")
	// fmt.Println(os.Hostname())
	fmt.Println(os.Getpid())

	return "create a file", fmt.Errorf(os.ErrClosed.Error())
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	config.ConnectRedis()
	config.ConnectDB()

	router.Run(":" + port)
}
