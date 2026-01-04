package controller

import (
	"strings"
	"time"

	"go-fiber-minimal/app/http/request"
	"go-fiber-minimal/app/http/resource"
	"go-fiber-minimal/database/entity"
	"go-fiber-minimal/lang"
	"go-fiber-minimal/lib"
	"go-fiber-minimal/util"

	"github.com/gofiber/fiber/v2"
)

var Auth authManager

type authManager struct{}

func (authManager) User(c *fiber.Ctx) error {
	user := c.Locals("user").(entity.User)
	roles := c.Locals("roles").([]string)
	permissions := c.Locals("permissions").([]string)

	id, _ := util.Hash.EncodeId(user.Id)
	return util.Api.SendData(
		c,
		lang.Trans.Convert(lang.Trans.Get().RETRIEVED_SUCCESSFULLY, fiber.Map{"operator": lang.Trans.Get().USER}),
		&resource.AuthShow{
			Id:          id,
			Username:    user.Username,
			Name:        user.Name,
			Roles:       roles,
			Permissions: permissions,
			CreatedAt:   util.Convert.Datetime2Str(user.CreatedAt),
			UpdatedAt:   util.Convert.Datetime2Str(user.UpdatedAt),
		},
	)
}

func (authManager) Login(c *fiber.Ctx) error {
	validated := &request.AuthLogin{}
	if err, isOk := lib.Validator.Check(c, validated); !isOk {
		return err
	}

	user := &entity.User{}
	if err := lib.DB.Run.Where(&entity.User{Username: validated.Username}).First(&user).Error; err != nil {
		return util.Api.SendError(c, lang.Trans.Convert(lang.Trans.Get().AUTH_FAILED))
	}

	if !util.Hash.BcryptVerify(validated.Password, user.Password) {
		return util.Api.SendError(c, lang.Trans.Convert(lang.Trans.Get().AUTH_FAILED))
	}

	hashId, err := util.Hash.EncodeId(user.Id)
	if err != nil {
		return util.Api.SendException(c, err)
	}

	token, err := lib.Jwt.Create(hashId)
	if err != nil {
		return util.Api.SendException(c, err)
	}

	auth := &entity.Auth{
		UserId: user.Id,
		Token:  token,
	}

	if err := lib.DB.Run.Create(auth).Error; err != nil {
		return util.Api.SendError(c, err.Error(), fiber.StatusInternalServerError)
	}

	return util.Api.SendCustom(c, resource.AuthLogin{
		Token: token,
	}, fiber.StatusOK)
}

func (authManager) Register(c *fiber.Ctx) error {
	validated := &request.AuthRegister{}
	if err, isOk := lib.Validator.Check(c, validated); !isOk {
		return err
	}

	tx := lib.DB.Run.Begin()

	user := &entity.User{}
	lib.DB.Run.Where(&entity.User{Username: validated.Username}).First(&user)
	if user.Id != 0 {
		return util.Api.SendError(c, lang.Trans.Convert(lang.Trans.Get().ALREADY_EXIST, fiber.Map{"operator": lang.Trans.Get().USER}))
	}

	newPassword, err := util.Hash.BcryptCreate(validated.Password)
	if err != nil {
		return util.Api.SendException(c, err)
	}

	user.Username = validated.Username
	user.Name = validated.Name
	user.Password = newPassword
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return util.Api.SendException(c, err)
	}

	roleUser := &entity.Role{}
	if err := tx.Model(&entity.Role{}).Where("name = ?", "user").First(roleUser).Error; err != nil {
		tx.Rollback()
		return util.Api.SendException(c, err)
	}

	uhr := &entity.UserHasRole{
		UserId: user.Id,
		RoleId: roleUser.Id,
	}
	if err := tx.Create(uhr).Error; err != nil {
		tx.Rollback()
		return util.Api.SendException(c, err)
	}

	tx.Commit()

	id, _ := util.Hash.EncodeId(user.Id)
	result := resource.AuthRegister{
		Data: resource.AuthShow{
			Id:        id,
			Username:  user.Username,
			Name:      user.Name,
			CreatedAt: util.Convert.Datetime2Str(user.CreatedAt),
			UpdatedAt: util.Convert.Datetime2Str(user.UpdatedAt),
		},
		Message: lang.Trans.Convert(lang.Trans.Get().SAVED_SUCCESSFULLY, fiber.Map{"operator": lang.Trans.Get().USER}),
	}

	return util.Api.SendCustom(c, result, fiber.StatusOK)
}

func (authManager) Logout(c *fiber.Ctx) error {
	getToken := c.Get("Authorization")
	token := ""
	if strings.HasPrefix(getToken, "Bearer ") {
		token = strings.TrimPrefix(getToken, "Bearer ")
	}

	if token != "" {
		if err := lib.DB.Run.Model(&entity.Auth{}).Where("token = ?", token).Updates(map[string]interface{}{
			"revoke":     true,
			"updated_at": time.Now(),
		}).Error; err != nil {
			return util.Api.SendException(c, err)
		}
	}

	return util.Api.SendSuccess(
		c,
		lang.Trans.Convert(lang.Trans.Get().LOGOUT_SUCCESSFULLY),
	)
}
