package handlers

import (
	"net/http"
	"text/template"

	model "Quester/model"

	"github.com/julienschmidt/httprouter"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	au := model.GetQuests()
	t := template.Must(template.ParseFiles("templates/index.html"))
	t.Execute(w, au)
}

// get [every unique npc, get every q of this npc], give this array to the template
func NpcHandlerGeneral(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	qgs := model.GetQuestgiversQualities()

	t := template.Must(template.ParseFiles("templates/npc.html"))
	t.Execute(w, qgs)
}

func NpcHandlerSpecial(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	questgiverQualities := model.GetQestgiverQualities()

	t := template.Must(template.ParseFiles("templates/npcQl.html"))
	t.Execute(w, questgiverQualities)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
