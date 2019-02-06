package repositories

import (
	"github.com/jinzhu/gorm"

	"github.com/pisatoo/pst-master/models"
)

type lxdRepo struct {
	DB *gorm.DB
}

type LxdRepository interface {
	Lxd(id int) (models.LXD, error)
	Lxds() ([]models.LXD, error)
	Create(l models.LXD) (models.LXD, error)
	Update(id int, l models.LXD) (models.LXD, error)
	Delete(id int) (bool, error)
}

func NewLxdRepo(DB *gorm.DB) LxdRepository {
	return &lxdRepo{DB}
}

func (r *lxdRepo) Lxd(id int) (models.LXD, error) {
	var l models.LXD

	r.DB.Find(&l, id)

	return l, nil
}

func (r *lxdRepo) Lxds() ([]models.LXD, error) {
	var lxds []models.LXD

	r.DB.Find(&lxds)

	return lxds, nil
}

func (r *lxdRepo) Create(l models.LXD) (models.LXD, error) {
	r.DB.Create(&l)

	return l, nil
}

func (r *lxdRepo) Update(id int, l models.LXD) (models.LXD, error) {
	var lxd models.LXD
	r.DB.First(&lxd, id)

	r.DB.Model(&lxd).Select([]string{"name", "updated_at"}).
		Updates(map[string]interface{}{"name": l.Name, "updated_at": l.UpdatedAt})

	return lxd, nil
}

func (r *lxdRepo) Delete(id int) (bool, error) {
	var lxd models.LXD
	r.DB.First(&lxd, id)

	r.DB.Delete(&lxd)

	return true, nil
}
