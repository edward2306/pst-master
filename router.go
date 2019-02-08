package main

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/pisatoo/pst-master/controllers"
	"github.com/pisatoo/pst-master/repositories"
)

func LoadRouter(db *gorm.DB) (r *mux.Router) {
	lxdRepo := repositories.NewLxdRepo(db)
	lxdController := controllers.NewLxdController(lxdRepo)
	lxcRepo := repositories.NewLxcRepo(db)
	lxcController := controllers.NewLxcController(lxcRepo)

	r = mux.NewRouter()
	v1 := r.PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/lxds", lxdController.Resources).Methods("GET", "POST")
	v1.HandleFunc("/lxds/{id}", lxdController.Resources).Methods("GET", "PATCH", "DELETE")
	v1.HandleFunc("/lxcs", lxcController.Resources).Methods("GET", "POST")
	v1.HandleFunc("/lxcs/{id}", lxcController.Resources).Methods("GET", "PATCH", "DELETE")

	return
}
