package authbyphone

import (
	"5-project/configs"
	"5-project/pkg/middleware"
	"5-project/pkg/req"
	"5-project/pkg/resp"
	"net/http"

	"github.com/gorilla/mux"
)

type AuthByPhoneHandlerDesp struct {
	*ServiceAuthByPhone
	*configs.Config
}

type AuthByPhoneHandler struct {
	*ServiceAuthByPhone
}

func NewHandlerAuthByPhone(mux *mux.Router, desp *AuthByPhoneHandlerDesp) {
	handler := AuthByPhoneHandler{
		ServiceAuthByPhone: desp.ServiceAuthByPhone,
	}
	mux.Handle("/authbyphone", middleware.Auth(handler.Auth(), desp.Config))
	mux.HandleFunc("/verifyByCode", handler.VerifyCode()).Methods("GET")
}

func (handler *AuthByPhoneHandler) Auth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RequestUserByPhone](w, r)
		if err != nil {
			resp.Json(w, err.Error(), 400)
			return
		}

		user, err := handler.ServiceAuthByPhone.AuthByPhone(body.Phone)
		if err != nil {
			resp.Json(w, err.Error(), 400)
			return
		}

		resp.Json(w, user, 200)
	}
}

func (handler *AuthByPhoneHandler) VerifyCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[ReqVerifyAuthByCode](w, r)
		if err != nil {
			resp.Json(w, err.Error(), 400)
			return
		}
		token, err := handler.ServiceAuthByPhone.VerifyByCode(body.SessionId, body.Code)
		if err != nil {
			resp.Json(w, err.Error(), 401) // 401 Unauthorized
			return
		}
		data := RespTokenAuthByCode{
			Token: token,
		}
		resp.Json(w, data, 200)
	}
}
