package model

type GetUserResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUserResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type DeleteUserRequest struct {
	Id string `json:"id"`
}
