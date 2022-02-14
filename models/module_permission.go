package models

import "time"

type ModulePermission struct {
	ID           int `gorm:"primary_key" json:"id"`
	ModuleAppId  int `gorm:"type:int;NOT NULL" json:"module_app_id" binding:"required"`
	ModuleApp    ModuleApp
	RoleId       int `gorm:"type:int;NOT NULL" json:"role_id" binding:"required"`
	Role         Role
	PermissionId int `gorm:"type:int;NOT NULL" json:"permission_id" binding:"required"`
	Permission   Permission
	Slug         string    `gorm:"type:varchar(100); NOT NULL" json:"slug"`
	CreatedAt    time.Time `json:"created_at"`
}

type ModulePermissions []ModulePermission
