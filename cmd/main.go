package main

import (
	// "GolangAdvanced/configs"
	"GolangAdvanced/configs"
	"GolangAdvanced/internal/auth"
	"GolangAdvanced/internal/link"
	"GolangAdvanced/internal/user"
	"GolangAdvanced/pkg/db"
	"GolangAdvanced/pkg/middleware"
	"context"
	"fmt"
	"net/http"
	"time"
)

func tickOperation(ctx context.Context) {
	ticker := time.NewTicker(200 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			fmt.Println("Tick")
		case <-ctx.Done():
			fmt.Println("Cancel")
			return
		}
	}
}


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
		Config: conf,
	})

	//Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	fmt.Println("Server is listening in port 8081")
	server.ListenAndServe()
}
