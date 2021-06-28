package routes

import (
	"go-curd-ops/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterUsersRoutes = func(router *mux.Router) {
	router.HandleFunc("/user/", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user/{UserId}", controllers.GetUserById).Methods("GET")

}