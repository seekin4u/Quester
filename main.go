package main

import (
	"Quester/handlers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	router.GET("/main", handlers.IndexHandler)
	router.GET("/npc", handlers.NpcHandlerGeneral)
	router.GET("/npc/1", handlers.NpcHandlerSpecial)
	http.ListenAndServe(":5000", router)

}
