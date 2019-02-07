package models

import (
	"time"
)

type LXC struct {
	ID        uint64    `gorm:"column:id; primary_key; AUTO_INCREMENT" json:"id"`
	Name      string    `gorm:"column:name; type:varchar(100)" json:"name"`
	IP        string    `gorm:"column:ip; type:varchar(15)" json:"ip"`
	VersionOs string    `gorm:"column:versionos; type:varchar(20)" json:"versionos`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
