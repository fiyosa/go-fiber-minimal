package seeder

import "gorm.io/gorm"

func Seed(g *gorm.DB) error {
	RolePermissionSeeder(g)
	UserSeeder(g)

	return nil
}
