package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()
	d := gin.Default()
	
	r.GET("/ping",r.HandleContext)
}
