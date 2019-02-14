package repositories_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/pisatoo/pst-master/models"
	LxcRepo "github.com/pisatoo/pst-master/repositories"

	"github.com/jinzhu/gorm"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetLxcById(t *testing.T) {
	db, mock, err := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

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

	rows := sqlmock.NewRows([]string{"id", "name", "ip", "os_version", "lxd_id", "created_at", "updated_at"}).
		AddRow(1, "amendo", "192.168.34.34", "Ubuntu 18.04", 1, now, now)

	query := "^SELECT (.+) FROM \"lxcs\" WHERE (.+)$"

	mock.ExpectQuery(query).WillReturnRows(rows)
	lr := LxcRepo.NewLxcRepo(gormDB)

	lxc := lr.Lxc(1)
	assert.NotNil(t, lxc)
	assert.Equal(t, l, lxc)
}

func TestGetAllLxc(t *testing.T) {
	db, mock, err := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	now := time.Now()

	rows := sqlmock.NewRows([]string{"id", "name", "ip", "os_version", "lxd_id", "created_at", "updated_at"}).
		AddRow(1, "amendo", "192.168.34.34", "Ubuntu 18.04", 1, now, now).
		AddRow(2, "mariesto", "192.168.26.26", "Ubuntu 16.04", 1, now, now)

	query := "^SELECT (.+) FROM \"lxcs\""

	mock.ExpectQuery(query).WillReturnRows(rows)
	lr := LxcRepo.NewLxcRepo(gormDB)

	lxcs := lr.Lxcs()
	assert.Len(t, lxcs, 2)
}

func TestCreateLxc(t *testing.T) {
	now := time.Now()

	l := models.LXC{
		Name:      "amendo",
		IP:        "192.168.34.34",
		OsVersion: "Ubuntu 18.04",
		IdLXD:     1,
		CreatedAt: now,
		UpdatedAt: now,
	}

	db, mock, err := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "^INSERT INTO \"lxcs\" (.+) VALUES (.+)"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(l.ID, l.Name, l.IP, l.CreatedAt, l.UpdatedAt).WillReturnResult(sqlmock.NewResult(1, 1))

	lr := LxcRepo.NewLxcRepo(gormDB)
	lxc := lr.Create(l)
	assert.NotNil(t, lxc)
}

func TestUpdateLxc(t *testing.T) {
	now := time.Now()

	l := models.LXC{
		Name:      "amendo",
		IP:        "192.168.34.1",
		OsVersion: "Ubuntu 16.04",
		IdLXD:     1,
		CreatedAt: now,
		UpdatedAt: now,
	}

	db, mock, err := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "^UPDATE \"lxcs\" set name=\\?, ip=\\?, os_version=\\?, lxd_id=\\?, updated_at=\\? WHERE id=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(l.Name, l.IP, l.UpdatedAt, l.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	lr := LxcRepo.NewLxcRepo(gormDB)
	lxc := lr.Update(1, l)
	assert.NotNil(t, lxc)
}

func TestDeleteLxc(t *testing.T) {
	db, mock, err := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "^DELETE FROM \"lxcs\" WHERE (.+)$"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	lr := LxcRepo.NewLxcRepo(gormDB)
	id := 1
	status := lr.Delete(id)
	assert.True(t, status)
}
