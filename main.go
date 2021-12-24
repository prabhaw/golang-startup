package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/prabhaw/golang-startup/config"
	"github.com/prabhaw/golang-startup/database"
	api_routes "github.com/prabhaw/golang-startup/routes"
)

func main(){
	serverPort := config.Config("PORT")
	err := database.Connect()
	if err != nil{
		log.Fatal(err)
		}
      app := fiber.New()
	  app.Use(cors.New())
	  app.Use(logger.New())
	  api_routes.RegisterApi(app)
      

	  err_lis := app.Listen(":"+ serverPort)
	  if err_lis != nil{
		 log.Fatal(err_lis) 
	  }
	  fmt.Println(fmt.Sprintf("Server is runing on PORT: %s", serverPort))
	}