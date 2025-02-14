package auth

type LoginReq struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Pin         string `json:"pin" validate:"required"`
}
