package entity

import (
	"time"
)

type RoleHasPermission struct {
	Id           int       `json:"id" gorm:"primaryKey;autoIncrement"`
	RoleId       int       `json:"role_id"`
	PermissionId int       `json:"permission_id"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:timestamp(0);not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"type:timestamp(0);not null;default:CURRENT_TIMESTAMP"`

	Role       Role       `gorm:"foreignKey:RoleId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Permission Permission `gorm:"foreignKey:PermissionId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
