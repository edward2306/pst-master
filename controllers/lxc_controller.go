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

type LxcController struct {
	repo repositories.LxcRepository
}

func NewLxcController(repo repositories.LxcRepository) LxcController {
	return LxcController{repo}
}

func (c *LxcController) Resources(w http.ResponseWriter, r *http.Request) {
	switch m := r.Method; m {
	case http.MethodGet:
		params := mux.Vars(r)
		if len(params) == 0 {
			c.Lxcs(w, r)
		} else {
			c.Lxc(w, r)
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

func (c *LxcController) Lxcs(w http.ResponseWriter, r *http.Request) {
	lxcs := c.repo.Lxcs()

	var cc []models.LXC

	for _, lxc := range lxcs {
		cc = append(cc, lxc)
	}

	respondWithJSON(w, http.StatusOK, cc)
}

func (c *LxcController) Lxc(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	lxc := c.repo.Lxc(id)

	respondWithJSON(w, http.StatusOK, lxc)
}

func (c *LxcController) Create(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	l := models.LXC{
		Name:      r.FormValue("name"),
		IP:        r.FormValue("ip"),
		CreatedAt: now,
		UpdatedAt: now,
	}

	lxc := c.repo.Create(l)

	respondWithJSON(w, http.StatusCreated, lxc)
}

func (c *LxcController) Update(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	l := models.LXC{
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

	lxc := c.repo.Update(id, l)

	respondWithJSON(w, http.StatusOK, lxc)
}

func (c *LxcController) Delete(w http.ResponseWriter, r *http.Request) {
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
