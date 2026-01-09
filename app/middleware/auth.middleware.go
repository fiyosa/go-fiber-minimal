package middleware

import (
	"go-fiber-minimal/database/entity"
	"go-fiber-minimal/lang"
	"go-fiber-minimal/lib"
	"go-fiber-minimal/service"
	"go-fiber-minimal/util"

	"github.com/gofiber/fiber/v2"
)

var (
	getPermissions *[]string
)

func Auth(permission ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := &entity.User{}
		if err, isOk := authentication(c, user); !isOk {
			return err
		}
		if err, isOk := authorization(c, permission...); !isOk {
			return err
		}
		return c.Next()
	}
}

func authentication(c *fiber.Ctx, user *entity.User) (error, bool) {
	// getToken := c.Get("Authorization")

	// if getToken == "" {
	// 	return util.Api.SendError(c, lang.Trans.Convert(lang.Trans.Get().UNAUTHORIZED_ACCESS)), false
	// }

	// tokenParts := strings.Split(getToken, " ")
	// if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
	// 	return util.Api.SendError(c, lang.Trans.Convert(lang.Trans.Get().UNAUTHORIZED_ACCESS)), false
	// }

	// token := tokenParts[1]
	// if _, err := lib.Jwt.Verify(token); err != nil {
	// 	return util.Api.SendError(c, lang.Trans.Convert(lang.Trans.Get().UNAUTHORIZED_ACCESS)), false
	// }

	token := Cookie.Get(c, "JWT")
	if token == "" {
		return util.Api.SendError(c, lang.Trans.Convert(lang.Trans.Get().UNAUTHORIZED_ACCESS)), false
	}

	auth := &entity.Auth{}
	if err := lib.DB.Run.Preload("User").Where(&entity.Auth{Token: token}).First(&auth).Error; err != nil {
		return util.Api.SendError(c, lang.Trans.Convert(lang.Trans.Get().UNAUTHORIZED_ACCESS)), false
	}
	if auth.Id == 0 {
		return util.Api.SendError(c, lang.Trans.Convert(lang.Trans.Get().UNAUTHORIZED_ACCESS)), false
	}

	roles := &[]string{}
	permissions := &[]string{}
	if err := service.RolePermission.GetRolesPermissionsByUserId(auth.UserId, roles, permissions); err != nil {
		return util.Api.SendError(c, err.Error()), false
	}

	getPermissions = permissions

	lib.LogFile.Info(fiber.Map{
		"user_name":   user.Username,
		"roles":       *roles,
		"permissions": *permissions,
	})

	*user = auth.User
	c.Locals("user", auth.User)
	c.Locals("roles", *roles)
	c.Locals("permissions", *permissions)
	c.Locals("token", token)
	return nil, true
}

func authorization(c *fiber.Ctx, permission ...string) (error, bool) {
	if len(permission) == 0 {
		return nil, true
	}

	check := false
	for _, v := range *getPermissions {
		if v == permission[0] {
			check = true
		}
	}

	if !check {
		return util.Api.SendError(c, lang.Trans.Convert(lang.Trans.Get().PERMISSION_FAILED)), false
	}

	return nil, true
}
