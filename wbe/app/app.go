package app

import (
	"devops/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Start(addr string) {
	r := gin.Default()

	user := r.Group("/api/user")
	role := r.Group("/api/role")
	setting := r.Group("/api/setting")
	dashboard := r.Group("/api/dashboard")
	host := r.Group("/api/host")
	cmd := r.Group("/api/cmd")
	monitor := r.Group("/api/monitor")
	audit := r.Group("/api/audit")

	{
		user.POST("/login", handler.UserLogin)
		user.POST("/create", handler.AuthMiddleware(), handler.CreateUser)
		user.POST("/getlist", handler.AuthMiddleware(), handler.ReadUserList)
		user.POST("/delete", handler.AuthMiddleware(), handler.DeleteUserByUserID)
		user.POST("/ResetPasswordByUserID", handler.AuthMiddleware(), handler.ResetPasswordByUserID)
		user.POST("/logout", handler.Logout)
		user.POST("/token2user", handler.AuthMiddleware(), handler.Token2User)
	}
	{
		role.POST("/create", handler.AuthMiddleware(), handler.CreateRole)
		role.POST("/getlist", handler.AuthMiddleware(), handler.ReadRoleList)
		role.POST("/delete", handler.AuthMiddleware(), handler.DeleteRoleByRoleID)
	}
	{
		setting.POST("/get", handler.AuthMiddleware(), handler.GetSetting)
		setting.POST("/set", handler.AuthMiddleware(), handler.SetSetting)
	}
	{
		dashboard.POST("/info", handler.AuthMiddleware(), handler.DashboardInfo)
	}
	{
		host.POST("/ReadHostList", handler.AuthMiddleware(), handler.ReadHostList)
		host.POST("/CreateHost", handler.AuthMiddleware(), handler.CreateHost)
		host.POST("/DeleteHost", handler.AuthMiddleware(), handler.DeleteHost)
		host.POST("/UpdateHost", handler.UpdateHost)
		host.POST("/ReadHost", handler.ReadHost)
	}
	{
		cmd.POST("/ExecCmd", handler.AuthMiddleware(), handler.ExecCmd)
	}
	{
		monitor.POST("/getcpu", handler.AuthMiddleware(), handler.GetCPU)
		monitor.POST("/getmemory", handler.AuthMiddleware(), handler.GetMemory)
	}
	{
		audit.POST("/log", handler.AuthMiddleware(), handler.GetLog)
	}

	err := r.Run(addr)
	handler.CheckError(err)
}
