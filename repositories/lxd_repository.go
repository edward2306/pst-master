package repositories

import (
	"github.com/jinzhu/gorm"

	"github.com/pisatoo/pst-master/models"
)

type lxdRepo struct {
	DB *gorm.DB
}

type LxdRepository interface {
	Lxd(id int) models.LXD
	Lxds() []models.LXD
	Create(l models.LXD) models.LXD
	Update(id int, l models.LXD) models.LXD
	Delete(id int) bool
}

func NewLxdRepo(DB *gorm.DB) LxdRepository {
	return &lxdRepo{DB}
}

func (r *lxdRepo) Lxd(id int) models.LXD {
	var l models.LXD

	r.DB.Find(&l, id)

	return l
}

func (r *lxdRepo) Lxds() []models.LXD {
	var lxds []models.LXD

	r.DB.Find(&lxds)

	return lxds
}

func (r *lxdRepo) Create(l models.LXD) models.LXD {
	r.DB.Create(&l)

	return l
}

func (r *lxdRepo) Update(id int, l models.LXD) models.LXD {
	var lxd models.LXD
	r.DB.First(&lxd, id)

	r.DB.Model(&lxd).Select([]string{"name", "updated_at"}).
		Updates(map[string]interface{}{"name": l.Name, "updated_at": l.UpdatedAt})

	return lxd
}

func (r *lxdRepo) Delete(id int) bool {
	var lxd models.LXD
	r.DB.First(&lxd, id)

	r.DB.Delete(&lxd)

	return true
}
