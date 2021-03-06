package bhginutils

import "github.com/gin-gonic/gin"

// Route : http route
type Route struct {
	HttpMethod   string            // HTTP method
	RelativePath string            // relative path
	HandlerFunc  []gin.HandlerFunc // handler func
}

// NewRoute new route
func NewRoute(httpMethod, relativePath string, controller gin.HandlerFunc, handlerFunc ...gin.HandlerFunc) *Route {
	return &Route{
		HttpMethod:   httpMethod,
		RelativePath: relativePath,
		HandlerFunc:  append(handlerFunc,controller),
	}
}

// RegisterRoutes : register routes
func RegisterRoutes(engine *gin.Engine, routes []*Route) {
	for _, route := range routes {
		engine.Handle(route.HttpMethod, route.RelativePath, route.HandlerFunc...)
	}
}

// RegisterRouteGroup : register route group
func RegisterRouteGroup(engine *gin.Engine, groupRelativePath string, routes []*Route) {
	group := engine.Group(groupRelativePath)

	for _, route := range routes {
		group.Handle(route.HttpMethod, route.RelativePath, route.HandlerFunc...)
	}
}


func RegistrGroup(engine *gin.Engine, groupRelativePath string) *gin.RouterGroup{
	return engine.Group(groupRelativePath)
}

func RegisterRouteByGroup(group *gin.RouterGroup,routes []*Route){
	for _, route := range routes {
		group.Handle(route.HttpMethod, route.RelativePath, route.HandlerFunc...)
	}
}
