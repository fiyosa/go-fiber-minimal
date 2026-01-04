package resource

type AuthShow struct {
	Id          string   `json:"id" example:""`
	Username    string   `json:"username" example:""`
	Name        string   `json:"name" example:""`
	Roles       []string `json:"roles" example:""`
	Permissions []string `json:"permissions" example:""`
	CreatedAt   string   `json:"created_at" example:""`
	UpdatedAt   string   `json:"updated_at" example:""`
}

type AuthLogin struct {
	Token string `json:"access_token" example:""`
}

type AuthRegister struct {
	Data    AuthShow `json:"data"`
	Message string   `json:"message" example:""`
}
