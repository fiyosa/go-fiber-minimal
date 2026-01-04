package entity

import (
	"time"
)

type RoleHasPermission struct {
	Id           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	RoleId       uint      `json:"role_id"`
	PermissionId uint      `json:"permission_id"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:timestamp(0);not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"type:timestamp(0);not null;default:CURRENT_TIMESTAMP"`

	Role       Role       `gorm:"foreignKey:RoleId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Permission Permission `gorm:"foreignKey:PermissionId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
