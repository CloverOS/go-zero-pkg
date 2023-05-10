package mysql_gorm

import (
	"gorm.io/gorm"
)

type GormBaseModel struct {
	CreatedAt int64          `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt int64          `json:"updated_at" gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type GormBaseModelWithId struct {
	Id uint `json:"id" gorm:"primaryKey"` //id
	GormBaseModel
}

type GormBaseModelWithOutLogicDelete struct {
	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime:milli"`
}
