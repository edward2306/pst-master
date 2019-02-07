package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/pisatoo/pst-master/models"
	"github.com/pisatoo/pst-master/repositories"
)

type LxdController struct {
	repo repositories.LxdRepository
}

func NewLxdController(repo repositories.LxdRepository) LxdController {
	return LxdController{repo}
}

func (c *LxdController) Resources(w http.ResponseWriter, r *http.Request) {
	switch m := r.Method; m {
	case http.MethodGet:
		params := mux.Vars(r)
		if len(params) == 0 {
			c.Lxds(w, r)
		} else {
			c.Lxd(w, r)
		}
	case http.MethodPost:
		c.Create(w, r)
	case http.MethodPatch:
		c.Update(w, r)
	case http.MethodDelete:
		c.Delete(w, r)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (c *LxdController) Lxds(w http.ResponseWriter, r *http.Request) {
	lxds := c.repo.Lxds()

	var dd []models.LXD

	for _, lxd := range lxds {
		dd = append(dd, lxd)
	}

	respondWithJSON(w, http.StatusOK, dd)
}

func (c *LxdController) Lxd(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	lxd := c.repo.Lxd(id)

	respondWithJSON(w, http.StatusOK, lxd)
}

func (c *LxdController) Create(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	l := models.LXD{
		Name:      r.FormValue("name"),
		IP:        r.FormValue("ip"),
		CreatedAt: now,
		UpdatedAt: now,
	}

	lxd := c.repo.Create(l)

	respondWithJSON(w, http.StatusCreated, lxd)
}

func (c *LxdController) Update(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	l := models.LXD{
		Name:      r.FormValue("name"),
		IP:        r.FormValue("ip"),
		UpdatedAt: now,
	}

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalln(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	lxd := c.repo.Update(id, l)

	respondWithJSON(w, http.StatusOK, lxd)
}

func (c *LxdController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ok := c.repo.Delete(id)

	if ok {
		respondWithJSON(w, http.StatusOK, true)
	}
}
