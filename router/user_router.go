package router

import "github.com/gin-gonic/gin"

func setupUserRouter(r *gin.Engine) {
	group := r.Group("/users")

	group.GET("/", func(r *gin.Context) {
		r.JSON(200, gin.H{
			"message": "pong",
		})

	})

}

func setupPingRouter(r *gin.Engine) {
	r.GET("/ping", func(r *gin.Context) {
		r.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
