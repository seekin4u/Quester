package main

import (
	"Quester/handlers"
	"Quester/model"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// TODO: add more handlers for read separate quest, delete quest.
// TODO: think of design of main page and how to quary information.
func main() {

	//model.GetQgs()

	/*fs := http.FileServer(http.Dir("static"))
	http.Handle("/style/", fs)

	http.HandleFunc("/", handlers.HandleJsonBig)
	http.HandleFunc("/main/", handlers.IndexHandler)
	http.HandleFunc("/npc/", handlers.NpcHandlerGeneral)
	//handlers.PrintNpc()

	log.Fatal(http.ListenAndServe(":5000", nil))*/
	model.GetQestgiverQualities()
	router := httprouter.New()
	router.GET("/main", handlers.IndexHandlerH)
	router.GET("/npc", handlers.NpcHandlerGeneralH)
	router.GET("/npc/1", handlers.NpcHandlerSpecial)
	http.ListenAndServe(":5000", router)

}
