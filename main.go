package main

import (
	"SysAuthenticate/config"
	"SysAuthenticate/database"
	"SysAuthenticate/internal/routes"
	"SysAuthenticate/logger"
	"github.com/gin-gonic/gin"
	"strconv"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	logger.Log.Infoln("Starting server...")
	configs, err := config.ConfigSet()
	if err != nil {
		panic(err)
	}
	db, err := database.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	serv := gin.Default()
	routes.RegisterSignupRoutes(serv, db)
	routes.RegisterUserRoutes(serv, db)
	routes.RegisterRoleRoutes(serv, db)
	routes.RegisterUserRoleRoutes(serv, db)
	routes.RegisterLoginRoutes(serv, db)
	routes.RegisterAuthRoutes(serv)
	_ = serv.Run(":" + strconv.Itoa(configs.Server.Port))
}
