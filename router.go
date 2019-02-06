package main

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/pisatoo/pst-master/controllers"
	"github.com/pisatoo/pst-master/repositories"
)

func LoadRouter(db *sql.DB) (r *mux.Router) {
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
