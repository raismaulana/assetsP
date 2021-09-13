package entity

import "time"

type Image struct {
	ID        int64     `gorm:"primary_key:auto_increment;column:id_image"`
	Title     string    `gorm:"type:varchar(255) not null"`
	Path      string    `gorm:"type:text not null"`
	CreatedAt time.Time ``
	UpdateAt  time.Time ``
}
