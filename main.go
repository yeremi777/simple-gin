package main

import (
	"fmt"
	"os"

	"simple-gin/configs"
	_ "simple-gin/migrations"
	"simple-gin/src/router"

	"github.com/gin-gonic/gin"
)

var validCommands = []string{
	"up",
	"down",
	"status",
	"reset",
	"up-by-one",
}

func isValidCommand(command string) bool {
	for _, validCommand := range validCommands {
		if command == validCommand {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) > 1 {
		if isValidCommand(os.Args[1]) {
			configs.LoadMigrations()
			return
		} else {
			return
		}
	}

	configs.LoadConfig()
	configs.LoadDB()

	r := gin.Default()
	api := r.Group("api")

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "OK"})
	})

	router.ProjectRouter(api.Group("/projects"))

	r.Run(fmt.Sprintf(":%v", configs.ENV.PORT))
}
