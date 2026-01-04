package seeder

import (
	"fmt"
	"go-fiber-minimal/database/entity"
	"os"

	"gorm.io/gorm"
)

var roles = []string{
	"admin",
	"user",
}

var permissions = []string{
	"user_index",
	"user_show",
}

func RolePermissionSeeder(g *gorm.DB) {
	tx := g.Begin()

	createRoles := []*entity.Role{}
	for _, v := range roles {
		createRoles = append(createRoles, &entity.Role{Name: v})
	}

	createPermissions := []*entity.Permission{}
	for _, v := range permissions {
		createPermissions = append(createPermissions, &entity.Permission{Name: v})
	}

	if err := tx.Create(&createRoles).Error; err != nil {
		tx.Rollback()
		fmt.Printf("Error seeder role: %v \n\n", err.Error())
		os.Exit(1)
	}
	if err := tx.Create(&createPermissions).Error; err != nil {
		tx.Rollback()
		fmt.Printf("Error seeder permission: %v \n\n", err.Error())
		os.Exit(1)
	}

	createRoleHasPermissions := []*entity.RoleHasPermission{}

	createRoleHasPermissions = append(createRoleHasPermissions, createAdmin(createRoles, createPermissions)...)
	createRoleHasPermissions = append(createRoleHasPermissions, createUser(createRoles, createPermissions)...)

	if err := tx.Create(&createRoleHasPermissions).Error; err != nil {
		tx.Rollback()
		fmt.Printf("Error seeder role has permission: %v \n\n", err.Error())
		os.Exit(1)
	}

	if err := tx.Commit().Error; err != nil {
		fmt.Println("Transaction commit failed role permission:", err)
		return
	}

	fmt.Println("Seeder: role permission created successfully.")
}

func createAdmin(cr []*entity.Role, cp []*entity.Permission) []*entity.RoleHasPermission {
	roleName := "admin"

	var roleID uint
	for _, v := range cr {
		if v.Name == roleName {
			roleID = v.Id
			break
		}
	}

	crhp := []*entity.RoleHasPermission{}
	for _, v := range cp {
		crhp = append(crhp, &entity.RoleHasPermission{
			RoleId:       roleID,
			PermissionId: v.Id,
		})
	}
	return crhp
}

func createUser(cr []*entity.Role, cp []*entity.Permission) []*entity.RoleHasPermission {
	roleName := "user"
	permissions := []string{
		"user_show",
	}

	var roleID uint
	for _, v := range cr {
		if v.Name == roleName {
			roleID = v.Id
			break
		}
	}

	crhp := []*entity.RoleHasPermission{}
	for _, v := range cp {
		for _, p := range permissions {
			if p == v.Name {
				crhp = append(crhp, &entity.RoleHasPermission{
					RoleId:       roleID,
					PermissionId: v.Id,
				})
			}
		}
	}
	return crhp
}
