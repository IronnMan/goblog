package bootstrap

import (
	"github.com/gorilla/mux"
	"goblog/pkg/route"
	routes2 "goblog/routes"
)

// SetupRoute 路由初始化
func SetupRoute() *mux.Router  {
	router := mux.NewRouter()
	routes2.RegisterWebRoutes(router)

	route.SetRoute(router)

	return router
}
