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

func NpcHandlerGeneral(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	qgs := model.GetQuestgiversQualities()

	t := template.Must(template.ParseFiles("templates/npc.html"))
	t.Execute(w, qgs)
}

func NpcHandlerSpecial(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	npc := ps.ByName("npc")
	questgiverQualitiesQuests := model.GetQestgiverQualitiesQuests(npc)

	t := template.Must(template.ParseFiles("templates/npcQl.html"))
	t.Execute(w, questgiverQualitiesQuests)
}

func QualitiesHandlerGeneral(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	qgs := model.GetQualityQuestgivers()

	t := template.Must(template.ParseFiles("templates/quality.html"))
	t.Execute(w, qgs)
}
