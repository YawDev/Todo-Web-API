package main

import (
	"os/exec"
	"runtime"
	app "todo-web-api/App"
	auth "todo-web-api/Authentication"
	"todo-web-api/sqlite_db"

	docs "todo-web-api/docs"

	gin "github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api/v1
// @title Todo Web API
// @version 1.0
// @description Todo Web API with JWT Auth
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apiKey JWT
// @in header
// @name token

func main() {
	sqlite_db.Connect()
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/")
		{
			RouteSetup(eg)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	go func() {
		err := openBrowser("http://localhost:8080/swagger/index.html")
		if err != nil {
			panic(err)
		}
	}()
	r.Run(":8080")
}

func RouteSetup(r *gin.RouterGroup) {
	r.POST("/Login", auth.Login)
	r.POST("/Register", auth.Register)
	r.GET("/GetUser/:id", auth.GetUserById)
	r.POST("/CreateList/:id", auth.AuthMiddleware(), app.CreateListForUser)
	r.DELETE("/DeleteList/:id", auth.AuthMiddleware(), app.DeleteList)
	r.GET("/GetList/:userid", app.GetListByUserId)
	r.POST("/CreateTask/:listid", auth.AuthMiddleware(), app.AddTaskToList)
	r.DELETE("/DeleteTask/:id", auth.AuthMiddleware(), app.DeleteTask)
	r.PUT("/UpdateTask/:id", auth.AuthMiddleware(), app.UpdateTask)
	r.PUT("/TaskCompleted/:id", auth.AuthMiddleware(), app.ChangeStatus)
	r.GET("/Home", auth.AuthMiddleware(), app.Home)
}

func openBrowser(url string) error {
	var err error
	switch runtime.GOOS {
	case "darwin": // macOS
		err = exec.Command("open", url).Start() // macOS-specific command
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	default:
		err = exec.Command("xdg-open", url).Start()
	}
	return err
}
