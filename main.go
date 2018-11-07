package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	route(r)
	r.Run(":8000")
}

func route(r *gin.Engine) {

	r.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message" : "pong",
		})
	})
	
}