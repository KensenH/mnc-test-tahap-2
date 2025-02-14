package payment

type CreatePaymentReq struct {
	Amount  float64 `json:"amount"`
	Remarks string  `json:"remarks"`
}
