package service

import "go-fiber-minimal/lib"

var RolePermission rolePermissionManager

type rolePermissionManager struct{}

func (rolePermissionManager) GetRolesByUserId(user_id uint, roles *[]string) error {
	return lib.DB.Run.
		Table("users AS u").
		Distinct("r.name").
		Joins("LEFT JOIN user_has_roles AS uhr ON uhr.user_id = u.id").
		Joins("LEFT JOIN roles AS r ON r.id = uhr.role_id").
		Where("u.id = ?", user_id).
		Debug().
		Scan(roles).
		Error
}

func (rolePermissionManager) GetPermissionsByRoles(roles []string, permissions *[]string) error {
	return lib.DB.Run.
		Table("permissions AS p").
		Distinct("p.name").
		Joins("LEFT JOIN role_has_permissions AS rhp ON rhp.permission_id = p.id").
		Joins("LEFT JOIN roles AS r ON r.id = rhp.role_id").
		Where("r.name IN ?", roles).
		Debug().
		Scan(permissions).
		Error
}

func (rolePermissionManager) GetPermissionsByUserId(user_id uint, permissions *[]string) error {
	return lib.DB.Run.
		Table("permissions AS p").
		Distinct("p.name").
		Joins("LEFT JOIN role_has_permissions AS rhp ON rhp.permission_id = p.id").
		Joins("LEFT JOIN roles AS r ON r.id = rhp.role_id").
		Joins("LEFT JOIN user_has_roles AS uhr ON uhr.role_id = r.id").
		Joins("LEFT JOIN users AS u ON u.id = uhr.user_id").
		Where("u.id = ?", user_id).
		Debug().
		Scan(permissions).
		Error
}

func (m rolePermissionManager) GetRolesPermissionsByUserId(user_id uint, roles *[]string, permissions *[]string) error {
	if err := m.GetRolesByUserId(user_id, roles); err != nil {
		return err
	}
	if err := m.GetPermissionsByUserId(user_id, permissions); err != nil {
		return err
	}
	return nil
}
