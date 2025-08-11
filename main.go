package main

import (
	"Quester/handlers"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	log.Println("Starting frontend")

	router := httprouter.New()

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, rec interface{}) {
		log.Printf("panic: %v", rec)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	// static + routes
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	router.GET("/main", handlers.IndexHandler)
	router.GET("/npc", handlers.NpcHandlerGeneral)
	router.GET("/npc/:npc", handlers.NpcHandlerSpecial)
	router.GET("/quality", handlers.QualitiesHandlerGeneral)
	router.GET("/quality/:quality", handlers.QualitiesHandlerSpecial)

	srv := &http.Server{
		Addr:         ":5000",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("Listening on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
	}()

	// Graceful shutdown on SIGINT/SIGTERM
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
	}
	log.Println("Bye.")
}
