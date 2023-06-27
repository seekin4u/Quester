package main

import (
	"log"
	"net/http"
	"os"

	handlers "Quester/handlers"
)

// TODO: add more handlers for read separate quest, delete quest.
// TODO: think of design of main page and how to quary information.
func main() {

	//styles loading
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/style/", fs)

	http.HandleFunc("/", handlers.HandleJson)
	http.HandleFunc("/main/", handlers.IndexHandler)

	log.Fatal(http.ListenAndServe(":5000", nil))

}

func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}
