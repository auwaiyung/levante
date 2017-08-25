package orm

import "time"

type OprBaseModel struct {
	ID        uint `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	CreatedBy string
	UpdatedBy string
	UpdatedAt time.Time
	IsDeleted bool
	DeletedAt *time.Time
}
