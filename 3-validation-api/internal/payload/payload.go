package payload

type LoginRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type DataRequest struct {
	Email string `json:"email"`
	Hash string  `json:"hash"`
}