package route

import "github.com/gin-gonic/gin"

type route func(engine *gin.Engine)
var routes = []route{}
var middlewares = []gin.HandlerFunc{}


func InitRoute() *gin.Engine{
	g := gin.Default()
	g.Use(middlewares...)

	for _,rt := range routes {
		rt(g)
	}
	return g
}
func RegisterRoute(route ...route) {
	routes = append(routes,route...)
}
func RegisterMiddleware(mid ...gin.HandlerFunc) {
	middlewares = append(middlewares,mid...)
}

