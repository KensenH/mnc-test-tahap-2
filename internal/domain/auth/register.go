package auth

type RegisterReq struct {
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Address     string `json:"address"`
	Pin         string `json:"pin" validate:"required"`
}
