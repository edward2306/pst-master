package main

import (
	"github.com/jinzhu/gorm"

	"github.com/gorilla/mux"
	"github.com/pisatoo/pst-master/controllers"
	"github.com/pisatoo/pst-master/repositories"
)

func LoadRouter(db *gorm.DB) (r *mux.Router) {
	lxdRepo := repositories.NewLxdRepo(db)
	lxdController := controllers.NewLxdController(lxdRepo)

	r = mux.NewRouter()
	r.HandleFunc("/lxds", lxdController.Lxds).Methods("GET")
	r.HandleFunc("/lxds", lxdController.Create).Methods("POST")
	r.HandleFunc("/lxds/{id}", lxdController.Lxd).Methods("GET")
	r.HandleFunc("/lxds/{id}", lxdController.Update).Methods("PATCH")
	r.HandleFunc("/lxds/{id}", lxdController.Delete).Methods("DELETE")

	return
}

func newRouter(db *gorm.DB) (r *mux.Router) {
	lxcRepo := repositories.NewLxcRepo(db)
	lxcController := controllers.NewLxcController(lxcRepo)

	r = mux.NewRouter()
	r.HandleFunc("/lxcs", lxcController.Lxcs).Methods("GET")
	r.HandleFunc("/lxcs", lxcController.Create).Methods("POST")
	r.HandleFunc("/lxcs/{id}", lxcController.Lxc).Methods("GET")
	r.HandleFunc("/lxcs/{id}", lxcController.Update).Methods("PATCH")
	r.HandleFunc("/lxcs/{id}", lxcController.Delete).Methods("DELETE")

	return
}
