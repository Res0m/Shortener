package main

import (
	// "GolangAdvanced/configs"
	"GolangAdvanced/configs"
	"GolangAdvanced/internal/auth"
	"GolangAdvanced/internal/link"
	"GolangAdvanced/internal/stat"
	"GolangAdvanced/internal/user"
	"GolangAdvanced/pkg/db"
	"GolangAdvanced/pkg/event"
	"GolangAdvanced/pkg/middleware"
	"fmt"
	"net/http"
)

// func tickOperation(ctx context.Context) {
// 	ticker := time.NewTicker(200 * time.Millisecond)
// 	for {
// 		select {
// 		case <-ticker.C:
// 			fmt.Println("Tick")
// 		case <-ctx.Done():
// 			fmt.Println("Cancel")
// 			return
// 		}
// 	}
// }

func App() http.Handler {
conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	//Repositories
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)

	//Services
	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})
	//Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		AuthService: authService,	
		Config:      conf,
	})

	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		EventBus:       eventBus,
		Config:         conf,
	})

	stat.NewStatHandler(router, stat.StatHandlerDeps{
		StatRepository: statRepository,
		Config:         conf,
	})
	
	go statService.AddClick()


	//Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)
	return stack(router)
}


func main() {
	
	app := App()
	server := http.Server{
		Addr:    ":8081",
		Handler: app,
	}


	fmt.Println("Server is listening in port 8081")
	server.ListenAndServe()
}
