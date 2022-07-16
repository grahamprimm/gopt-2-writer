package main

import (
	"fmt"
	"os"
	//"github.com/grahamprimm/gopt-2-writer/log"
)

func main() {
	//err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	app := fiber.New()

	//app.Use(csrf.New())
	app.Use(logger.New())

	//setupRoutes(app)

	sugarD(app.Listen(os.Getenv("APP_PORT")))
}
