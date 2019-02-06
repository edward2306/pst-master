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

func modifyResponse(lxd models.LXD) map[string]interface{} {
	l := make(map[string]interface{})
	l["id"] = lxd.ID
	l["name"] = lxd.Name
	l["ip"] = lxd.IP
	l["created_at"] = lxd.CreatedAt
	l["updated_at"] = lxd.UpdatedAt

	return l
}

func (c *LxdController) Lxds(w http.ResponseWriter, r *http.Request) {
	lxd, err := c.repo.Lxds()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, lxd)
}

func (c *LxdController) Lxd(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])
	lxd, err := c.repo.Lxd(id)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

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

	res, err := c.repo.Create(l)
	lxd := modifyResponse(res)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

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

	res, err := c.repo.Update(id, l)
	lxd := modifyResponse(res)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, lxd)
}

func (c *LxdController) Delete(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ok, err := c.repo.Delete(id)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if ok {
		respondWithJSON(w, http.StatusOK, true)
	}
}
