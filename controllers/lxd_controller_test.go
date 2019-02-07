package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"

	"github.com/pisatoo/pst-master/controllers"

	"github.com/pisatoo/pst-master/repositories/mocks"

	"github.com/pisatoo/pst-master/models"
)

func TestLxds(t *testing.T) {
	ll := make([]models.LXD, 1)
	mockLxdRepo := new(mocks.LxdRepository)
	mockLxdRepo.On("Lxds").Return(ll, nil)

	lc := controllers.NewLxdController(mockLxdRepo)

	req, err := http.NewRequest("GET", "/lxds", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/lxds", lc.Lxds)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Controller returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestLxd(t *testing.T) {
	now := time.Now()

	l := models.LXD{
		ID:        0,
		Name:      "amendo",
		IP:        "192.168.34.34",
		CreatedAt: now,
		UpdatedAt: now,
	}

	id := int(l.ID)

	mockLxdRepo := new(mocks.LxdRepository)
	mockLxdRepo.On("Lxd", id).Return(l, nil)

	lc := controllers.NewLxdController(mockLxdRepo)

	req, err := http.NewRequest("GET", "/lxds/"+strconv.Itoa(id), nil)
	if err != nil {
		t.Fatalf(err.Error())
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/lxds/{id}", lc.Lxd)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Controller returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestCreate(t *testing.T) {
	now := time.Now()

	l := models.LXD{
		Name:      "amendo",
		IP:        "192.168.34.34",
		CreatedAt: now,
		UpdatedAt: now,
	}

	tempLxd := l
	tempLxd.ID = 1

	mockLxdRepo := new(mocks.LxdRepository)
	mockLxdRepo.On("Create", mock.AnythingOfType("models.LXD")).Return(l, nil)

	lc := controllers.NewLxdController(mockLxdRepo)

	payload, err := json.Marshal(tempLxd)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/lxds", strings.NewReader(string(payload)))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/lxds", lc.Create)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Controller returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestUpdate(t *testing.T) {
	now := time.Now()

	l := models.LXD{
		ID:        1,
		Name:      "amendo",
		IP:        "192.168.34.34",
		CreatedAt: now,
		UpdatedAt: now,
	}

	id := int(l.ID)

	mockLxdRepo := new(mocks.LxdRepository)
	mockLxdRepo.On("Update", id, mock.AnythingOfType("models.LXD")).Return(l, nil)

	lc := controllers.NewLxdController(mockLxdRepo)

	payload, err := json.Marshal(l)
	assert.NoError(t, err)

	req, err := http.NewRequest("PATCH", "/lxds/"+strconv.Itoa(id), strings.NewReader(string(payload)))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/lxds/{id}", lc.Update)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Controller returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestDetele(t *testing.T) {
	now := time.Now()

	l := models.LXD{
		ID:        1,
		Name:      "amendo",
		IP:        "192.168.34.34",
		CreatedAt: now,
		UpdatedAt: now,
	}

	id := int(l.ID)

	mockLxdRepo := new(mocks.LxdRepository)
	mockLxdRepo.On("Delete", id).Return(true, nil)

	lc := controllers.NewLxdController(mockLxdRepo)

	req, err := http.NewRequest("DELETE", "/lxds/"+strconv.Itoa(id), nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/lxds/{id}", lc.Delete)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Controller returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}
