package router

import (

	"net/http"
	"bird_golang_back/internal/app"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func RouterEngine() *gin.Engine {
	r := gin.Default()

	r.Use(CORS())
	r.POST("/update-bird-data",app.UpdateBirdData)
	r.GET("/get-bird-data",app.GetBirdData)
	r.POST("/update-bird-des",app.UpdateBirdDes)

	
	r.StaticFS("/download", http.Dir("/bird_image"))

	return r
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		targets := viper.GetString("cors.target")
		c.Writer.Header().Set("Access-Control-Allow-Origin", targets)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
