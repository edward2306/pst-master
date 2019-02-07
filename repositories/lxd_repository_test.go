package repositories_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/pisatoo/pst-master/models"
	LxdRepo "github.com/pisatoo/pst-master/repositories"

	"github.com/jinzhu/gorm"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	now := time.Now()

	l := models.LXD{
		ID:        1,
		Name:      "amendo",
		IP:        "192.168.34.34",
		CreatedAt: now,
		UpdatedAt: now,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "ip", "created_at", "updated_at"}).
		AddRow(1, "amendo", "192.168.34.34", now, now)

	query := "^SELECT (.+) FROM \"lxds\" WHERE (.+)$"

	mock.ExpectQuery(query).WillReturnRows(rows)
	lr := LxdRepo.NewLxdRepo(gormDB)

	lxd := lr.Lxd(1)
	assert.NotNil(t, lxd)
	assert.Equal(t, l, lxd)
}

func TestGetAllLxd(t *testing.T) {
	db, mock, err := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	now := time.Now()

	rows := sqlmock.NewRows([]string{"id", "name", "ip", "created_at", "updated_at"}).
		AddRow(1, "amendo", "192.168.34.34", now, now).
		AddRow(2, "mariesto", "192.168.26.26", now, now)

	query := "^SELECT (.+) FROM \"lxds\""

	mock.ExpectQuery(query).WillReturnRows(rows)
	lr := LxdRepo.NewLxdRepo(gormDB)

	lxds := lr.Lxds()
	assert.Len(t, lxds, 2)
}

func TestCreateLxd(t *testing.T) {
	now := time.Now()

	l := models.LXD{
		Name:      "amendo",
		IP:        "192.168.34.34",
		CreatedAt: now,
		UpdatedAt: now,
	}

	db, mock, err := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "^INSERT INTO \"lxds\" (.+) VALUES (.+)"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(l.ID, l.Name, l.IP, l.CreatedAt, l.UpdatedAt).WillReturnResult(sqlmock.NewResult(1, 1))

	lr := LxdRepo.NewLxdRepo(gormDB)
	lxd := lr.Create(l)
	assert.NotNil(t, lxd)
}

func TestUpdateLxd(t *testing.T) {
	now := time.Now()

	l := models.LXD{
		Name:      "amendo",
		IP:        "192.168.34.34",
		UpdatedAt: now,
	}

	db, mock, err := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "^UPDATE \"lxds\" set name=\\?, ip=\\?, updated_at=\\? WHERE id=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(l.Name, l.IP, l.UpdatedAt, l.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	lr := LxdRepo.NewLxdRepo(gormDB)
	lxd := lr.Update(1, l)
	assert.NotNil(t, lxd)
}

func TestDeleteLxd(t *testing.T) {
	db, mock, err := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "^DELETE FROM \"lxds\" WHERE (.+)$"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	lr := LxdRepo.NewLxdRepo(gormDB)
	id := 1
	status := lr.Delete(id)
	assert.True(t, status)
}
