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

func TestLxcs(t *testing.T) {
	ll := make([]models.LXC, 1)
	mockLxcRepo := new(mocks.LxcRepository)
	mockLxcRepo.On("Lxcs").Return(ll)

	lc := controllers.NewLxcController(mockLxcRepo)

	req, err := http.NewRequest("GET", "/Lxcs", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/Lxcs", lc.Lxcs)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Controller returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestLxc(t *testing.T) {
	now := time.Now()

	l := models.LXC{
		ID:        1,
		Name:      "amendo",
		IP:        "192.168.34.34",
		OsVersion: "Ubuntu 18.04",
		IdLXD:     1,
		CreatedAt: now,
		UpdatedAt: now,
	}

	id := int(l.ID)

	mockLxcRepo := new(mocks.LxcRepository)
	mockLxcRepo.On("Lxc", id).Return(l)

	lc := controllers.NewLxcController(mockLxcRepo)

	req, err := http.NewRequest("GET", "/Lxcs/"+strconv.Itoa(id), nil)
	if err != nil {
		t.Fatalf(err.Error())
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/Lxcs/{id}", lc.Lxc)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Controller returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestCreateLxc(t *testing.T) {
	now := time.Now()

	l := models.LXC{
		ID:        1,
		Name:      "amendo",
		IP:        "192.168.34.34",
		OsVersion: "Ubuntu 18.04",
		IdLXD:     1,
		CreatedAt: now,
		UpdatedAt: now,
	}

	tempLxc := l
	tempLxc.ID = 1

	mockLxcRepo := new(mocks.LxcRepository)
	mockLxcRepo.On("Create", mock.AnythingOfType("models.LXC")).Return(l)

	lc := controllers.NewLxcController(mockLxcRepo)

	payload, err := json.Marshal(tempLxc)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/Lxcs", strings.NewReader(string(payload)))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/Lxcs", lc.Create)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Controller returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestUpdateLxc(t *testing.T) {
	now := time.Now()

	l := models.LXC{
		ID:        1,
		Name:      "Edward",
		IP:        "192.168.34.184",
		OsVersion: "Ubuntu 18.04",
		IdLXD:     1,
		CreatedAt: now,
		UpdatedAt: now,
	}

	id := int(l.ID)

	mockLxcRepo := new(mocks.LxcRepository)
	mockLxcRepo.On("Update", id, mock.AnythingOfType("models.LXC")).Return(l)

	lc := controllers.NewLxcController(mockLxcRepo)

	payload, err := json.Marshal(l)
	assert.NoError(t, err)

	req, err := http.NewRequest("PATCH", "/Lxcs/"+strconv.Itoa(id), strings.NewReader(string(payload)))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/Lxcs/{id}", lc.Update)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Controller returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestDeteleLxc(t *testing.T) {
	now := time.Now()

	l := models.LXC{
		ID:        1,
		Name:      "amendo",
		IP:        "192.168.34.34",
		OsVersion: "Ubuntu 18.04",
		IdLXD:     1,
		CreatedAt: now,
		UpdatedAt: now,
	}

	id := int(l.ID)

	mockLxcRepo := new(mocks.LxcRepository)
	mockLxcRepo.On("Delete", id).Return(true)

	lc := controllers.NewLxcController(mockLxcRepo)

	req, err := http.NewRequest("DELETE", "/Lxcs/"+strconv.Itoa(id), nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/Lxcs/{id}", lc.Delete)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Controller returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}
