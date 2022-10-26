package controller

import (
	Application "auth/application"
	DomainInputPort "auth/domain/input-port"
	"encoding/json"
	"net/http"
)

type AuthController struct {
	muxServer   *http.ServeMux
	application DomainInputPort.AuthInputPort
}

func NewAuthController(muxServer *http.ServeMux) *AuthController {
	application := Application.NewAuthApplication()
	return &AuthController{muxServer, application}
}

func (a *AuthController) RegisterRoutes() {
	prefix := "/auth"
	a.muxServer.HandleFunc(prefix+"/login", a.Login)
	a.muxServer.HandleFunc(prefix+"/signup", a.SignUp)
}

func (a *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	user, err := a.application.Login("email", "password")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func (a *AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	user, err := a.application.SignUp("email", "fullname", "password")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
