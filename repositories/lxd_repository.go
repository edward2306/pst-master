package repositories

import (
	"database/sql"
	"log"

	"github.com/pisatoo/pst-master/models"
)

type lxdRepo struct {
	DB *sql.DB
}

type LxdRepository interface {
	Lxd(id int) (models.LXD, error)
	Lxds() ([]models.LXD, error)
	Create(u models.LXD) (models.LXD, error)
	Update(id int, u models.LXD) (models.LXD, error)
	Delete(id int) (bool, error)
}

func NewLxdRepo(DB *sql.DB) LxdRepository {
	return &lxdRepo{DB}
}

func (r *lxdRepo) Lxd(id int) (models.LXD, error) {
	var u models.LXD

	err := r.DB.QueryRow("select * from lxd where id=?", id).
		Scan(&u.ID, &u.Name, &u.IP, &u.CreatedAt, &u.UpdatedAt)

	return u, err
}

func (r *lxdRepo) Lxds() ([]models.LXD, error) {
	var lxds []models.LXD

	rows, err := r.DB.Query(`select * from lxd`)
	if err != nil {
		log.Fatalln(err)
	}

	defer rows.Close()

	for rows.Next() {
		var u models.LXD
		err := rows.Scan(&u.ID, &u.Name, &u.IP, &u.CreatedAt, &u.UpdatedAt)

		if err != nil {
			log.Fatalln(err)
		}

		lxds = append(lxds, u)
	}

	return lxds, nil
}

func (r *lxdRepo) Create(u models.LXD) (models.LXD, error) {
	stmt, err := r.DB.Prepare("insert lxd set name=?, ip=?, created_at=?, updated_at=?")
	if err != nil {
		log.Fatalln(err)
	}

	res, err := stmt.Exec(u.Name, u.IP, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		log.Fatalln(err)
	}

	id, err := res.LastInsertId()
	idInt := int(id)

	lxd, _ := r.Lxd(idInt)

	return lxd, err
}

func (r *lxdRepo) Update(i int, u models.LXD) (models.LXD, error) {
	stmt, err := r.DB.Prepare("update lxd set name=?, ip=?, updated_at=? where id=?")

	if err != nil {
		log.Fatalln(err)
	}

	_, err = stmt.Exec(u.Name, u.IP, u.UpdatedAt, i)
	if err != nil {
		log.Fatalln(err)
	}

	lxd, _ := r.Lxd(i)

	return lxd, err
}

func (r *lxdRepo) Delete(id int) (bool, error) {
	stmt, err := r.DB.Prepare("delete from lxd where id=?")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatalln(err)
	}

	return true, nil
}
