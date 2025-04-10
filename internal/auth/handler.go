package auth

import (
	"GolangAdvanced/configs"
	res "GolangAdvanced/pkg/response"
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (login *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println(login.Config.Auth.Secret)
		fmt.Println("Login")
		//Прочитать Body
		var payload LoginRequest
		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			res.JsonRes(w, err.Error(), 402)
			return
		}
		if payload.Email == "" {
			res.JsonRes(w, "Email required", 402)
			return
		}
		_, err = mail.ParseAddress(payload.Email)
		if err != nil {
			res.JsonRes(w, "Wrond Email", 402)
			return
		}
		// match, _ := regexp.MatchString(`[A-Za-z0-9\._%+\-]+@[A-Za-z0-9\.\-]+\.[A-Za-z]{2,}`,payload.Email)
		// if match{
		// 	res.JsonRes(w, "Wrong Email", 402)
		// 	return
		// }
		if payload.Password == "" {
			res.JsonRes(w, "Password required", 402)
			return
		}
		fmt.Println(payload)
		data := LoginResponse{
			Token: "123",
		}
		res.JsonRes(w, data, 200)
	}
}

func (reg *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Register")
	}
}
