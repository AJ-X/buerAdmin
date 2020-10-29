package router

import (
	"buerAdmin/controller/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "welcome")
	})
	api := r.Group("/api")
	{
		api.GET("/login", func(c *gin.Context) {
			user := auth.Buer_admin{}
			user.Login(c)
		})
	}
	_ = r.Run(":8888")
}
