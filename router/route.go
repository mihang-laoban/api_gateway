package router

import "github.com/gin-gonic/gin"

var RouteMapping = make(map[string]gin.HandlerFunc)

func init() {
	RouteMapping["ping"] = Ping
}
