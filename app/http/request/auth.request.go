package request

type AuthLogin struct {
	Username string `json:"username" form:"username" validate:"required,min=3" example:""`
	Password string `json:"password" form:"password" validate:"required,min=3" example:""`
}

type AuthRegister struct {
	Username string `json:"username" form:"username" from:"username" validate:"required,min=3" example:""`
	Password string `json:"password" form:"password" from:"password" validate:"required,min=3" example:""`
	Name     string `json:"name" form:"name" from:"name" validate:"required,min=3" example:""`
}
