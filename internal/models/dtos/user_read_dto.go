package dtos

type UserReadDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
