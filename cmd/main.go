package main

import (
	// "GolangAdvanced/configs"
	"GolangAdvanced/configs"
	"GolangAdvanced/internal/auth"
	"GolangAdvanced/internal/link"
	"GolangAdvanced/internal/user"
	"GolangAdvanced/pkg/db"
	"GolangAdvanced/pkg/middleware"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	//Repositories
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)

	//Services
	authService := auth.NewAuthService(userRepository)

	//Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})

	link.NewLinkHandler(router, &link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: middleware.Logging(router),
	}

	fmt.Println("Server is listening in port 8081")
	server.ListenAndServe()
}
