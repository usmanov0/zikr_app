package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
	"zikar-app/internal/adapter"
	"zikar-app/internal/app"
	"zikar-app/internal/common"
	handler "zikar-app/internal/port/http"
)

func main() {
	httpServer()
}

func httpServer() *chi.Mux {
	db, err := common.ConnectToDb(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
	)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer db.Close()

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Hello, Chi!"))
	})

	prayRepo := adapter.NewPrayRepo(db)
	prayUseCase := app.NewBasketService(prayRepo)
	prayHandler := handler.NewPrayHandler(prayUseCase)

	// Routers
	router.Route("/api", func(r chi.Router) {

		r.Route("/pray", func(r chi.Router) {
			r.Post("/save", prayHandler.PraySave)
			r.Post("/read", prayHandler.PrayRead)
			r.Post("/put", prayHandler.PrayUpdate)
			r.Post("/remove", prayHandler.PrayDelete)
		})
	})

	server := &http.Server{Addr: os.Getenv("HTTP_PORT"), Handler: router}
	log.Println("Starting server on port...", os.Getenv("HTTP_PORT"))
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		panic(err)
	}
	defer server.Close()

	return router
}
