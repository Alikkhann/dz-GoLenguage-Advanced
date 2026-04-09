package authbyphone

type RequestUserByPhone struct {
	Phone string `json:"phone" validate:"required"`
}

type ResponseUserByPhone struct {
	SessionId string `json:"sessionId"`
	Code      int    `json:"code"`
}

type ReqVerifyAuthByCode struct {
	SessionId string `json:"sessionId" validate:"required"`
	Code      int    `json:"code" validate:"required"`
}

type RespTokenAuthByCode struct {
	Token string `json:"token" validate:"required"`
}









// package authbyphone

// type RequestUserByPhone struct {
// 	Email    string `json:"email" validate:"required,email"`
// 	Password string `json:"password" validate:"required"`
// 	Name     string `json:"name" validate:"required"`
// 	Phone    string `json:"phone" validate:"required"`
// }

// type ResponseUserAndPhone struct {
// 	SessionId string `json:"sessionId"`
// 	Code      int    `json:"code"`
// }
