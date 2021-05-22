package public

import (
	"api_gateway/conf"
	"api_gateway/router"
	"dev-go-base/server"
	"fmt"
	"github.com/gin-gonic/gin"
)

var App *Application

type (
	Application struct {
		Engine  *gin.Engine
		AppConf *AppConf
	}

	AppConf struct {
		Port int
	}
)

func (app *Application) Init() {
	App = &Application{
		server.NewServer(),
		&AppConf{Port: conf.C.Base.Port},
	}
	App.InitRouter()
}

func (app *Application) InitRouter() {
	v1 := App.Engine.Group("/v1")
	{
		for k, v := range router.RouteMapping {
			v1.POST(k, v)
		}
	}
}

func Run() {
	addr := fmt.Sprintf(":%d", App.AppConf.Port)
	if err := App.Engine.Run(addr); err != nil {
		panic(err)
	}
}
