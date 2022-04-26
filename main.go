package main

import (
	"GlassKT/database"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()

	g := gin.Default()

	g.Run()

}
