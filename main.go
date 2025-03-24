package main

import (
	"SysAuthenticate/database"
	"SysAuthenticate/internal/routes"
	"github.com/gin-gonic/gin"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
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

	_ = serv.Run(":8080")
}
