package util

import "slices"

var Role roleManager

type roleManager struct{}

func checkRole(check string, roles []string) bool {
	return slices.Contains(roles, check)
}

// ["admin"] => true
func (roleManager) IsAdmin(roles []string) bool {
	return checkRole("admin", roles)
}

// ["user"] => true
func (roleManager) IsUser(roles []string) bool {
	return checkRole("user", roles)
}
