package routes

import (
	"github.com/gorilla/mux"
)

// InitRoutes inicia las rutas
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	SetLoginRouter(router)
	SetUserRouter(router)
	SetCommentRouter(router)
	SetVoteRouter(router)
	SetRealtimeRouter(router)
	SetPublicRouter(router)

	return router
}
