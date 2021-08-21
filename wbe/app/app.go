package app

import (
	"devops/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Start(addr string) {
	r := gin.Default()
	user := r.Group("/api/user")
	host := r.Group("/api/host")
	cmd := r.Group("/api/cmd")
	host.Use(handler.AuthMiddleware())
	cmd.Use(handler.AuthMiddleware())
	{
		user.POST("/login", handler.UserLogin)
		user.POST("/logout", handler.Logout)
		user.POST("/token2user", handler.AuthMiddleware(), handler.Token2User)
	}
	{
		host.POST("/CreateHost", handler.CreateHost)
		host.POST("/UpdateHost", handler.UpdateHost)
		host.POST("/ReadHostList", handler.ReadHostList)
		host.POST("/ReadHost", handler.ReadHost)
		host.POST("/DeleteHost", handler.DeleteHost)
	}
	{
		cmd.POST("/ExecCmd", handler.ExecCmd)
	}
	err := r.Run(addr)
	handler.CheckError(err)
}
