package api_routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prabhaw/golang-startup/module/user"
)

func RegisterApi(app *fiber.App){
   api:= app.Group("/api")
  
    user.UserRouter(api.Group("/user"))

}