package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/pisatoo/pst-master/models"
)

type lxcRepo struct {
	DB *gorm.DB
}

type LxcRepository interface {
	Lxc(id int) models.LXC
	Lxcs() []models.LXC
	Create(l models.LXC) models.LXC
	Update(id int, l models.LXC) models.LXC
	Delete(id int) bool
}

func NewLxcRepo(DB *gorm.DB) LxcRepository {
	return &lxcRepo{DB}
}

func (r *lxcRepo) Lxc(id int) models.LXC {
	var l models.LXC

	r.DB.Find(&l, id)

	return l
}

func (r *lxcRepo) Lxcs() []models.LXC {
	var ll []models.LXC

	r.DB.Find(&ll)

	return ll
}

func (r *lxcRepo) Create(l models.LXC) models.LXC {
	r.DB.Create(&l)

	return l
}

func (r *lxcRepo) Update(id int, l models.LXC) models.LXC {
	var lxc models.LXC
	r.DB.First(&lxc, id)

	r.DB.Model(&lxc).Select([]string{"name", "updated_at"}).
		Updates(map[string]interface{}{"name": l.Name, "updated_at": l.UpdatedAt})

	return lxc
}

func (r *lxcRepo) Delete(id int) bool {
	var lxc models.LXC
	r.DB.First(&lxc, id)

	r.DB.Delete(&lxc)

	return true
}
