package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "zeus/docs"
	"zeus/pkg/api/controllers"
	"zeus/pkg/api/middleware"
)

func Init(e *gin.Engine) {
	e.Use(
		gin.Recovery(),
	)
	e.Use(cors.Default()) // CORS
	e.Use(middleware.SetLangVer())
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	e.GET("/test", controllers.Healthy)
	//version fragment
	v1 := e.Group("/v1")
	jwtAuth := middleware.JwtAuth()
	//auth.POST("/token", jwtAuth.LoginHandler)
	//auth.GET("/refresh_token", jwtAuth.RefreshHandler)


	//api handlers
	v1.POST("/users/login",jwtAuth.LoginHandler)
	v1.POST("/users/login/refresh",jwtAuth.RefreshHandler)

	v1.Use(jwtAuth.MiddlewareFunc(), middleware.JwtPrepare)
	userController := controllers.UserController{}

	v1.GET("/login/info", userController.Info)
	//update login user's password
	v1.PUT("/login/password",userController.EditLoginUserPassword)
	//user
	v1.GET("/users", userController.List)
	v1.GET("/users/:id", userController.Get)
	v1.PUT("/users/:id", userController.Edit)
	v1.PUT("/users/:id/status", userController.EditStatus)
	v1.PUT("/users/:id/password", userController.EditPassword)
	v1.DELETE("/users/:id", userController.Delete)


	roleController := controllers.RoleController{}
	//role
	v1.GET("/roles", roleController.List)
	v1.GET("/roles/:id", roleController.Get)
}
