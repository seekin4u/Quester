package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"

	model "Quester/model"

	"github.com/julienschmidt/httprouter"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	au := model.ShowAllQuests()
	t := template.Must(template.ParseFiles("templates/index.html"))
	t.Execute(w, au)
}

func IndexHandlerH(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	au := model.ShowAllQuests()
	t := template.Must(template.ParseFiles("templates/index.html"))
	t.Execute(w, au)
}

func NpcHandlerGeneralH(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	au := model.ShowAllQuests()

	var npcs []string
	for _, el := range au.Quests {
		if !contains(npcs, el.Quest.QuestReward.QuestgiverName) {
			npcs = append(npcs, el.Quest.QuestReward.QuestgiverName)
		}
	}

	var qgs model.Questgivers
	qgs.Questgivers = npcs

	t := template.Must(template.ParseFiles("templates/npc.html"))
	t.Execute(w, qgs)
}

func NpcHandlerSpecial(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	questgiverQualities := model.GetQestgiverQualities()

	t := template.Must(template.ParseFiles("templates/npcQl.html"))
	t.Execute(w, questgiverQualities)
}

func NpcHandlerGeneral(w http.ResponseWriter, r *http.Request) {
	au := model.ShowAllQuests()

	t := template.Must(template.ParseFiles("templates/npc.html"))
	t.Execute(w, au)
}

func NpcHandler(w http.ResponseWriter, r *http.Request, parameter string) {
	au := model.ShowAllQuests()

	t := template.Must(template.ParseFiles("templates/npc.html"))
	t.Execute(w, au)
}

func HandleJson(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var receivedQuest model.QuestStructure
	err = json.Unmarshal(body, &receivedQuest)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)

	var questTime model.QuestTime
	questTime.Time = time.Now().Unix()
	questTime.Quest = receivedQuest

	fmt.Println("---------")

	fmt.Println("Time is " + string(strconv.FormatInt(questTime.Time, 10)))
	fmt.Println("	 Character:" + questTime.Quest.Character)
	fmt.Println("	 Content:" + questTime.Quest.Content)
	if len(questTime.Quest.QuestReward.QuestgiverName) != 0 {
		fmt.Println("	 	QG:" + questTime.Quest.QuestReward.QuestgiverName)
	}
	if len(questTime.Quest.QuestReward.RewardLp) != 0 {
		fmt.Println("		LP:" + questTime.Quest.QuestReward.RewardLp)
	}
	if len(questTime.Quest.QuestReward.RewardExp) != 0 {
		fmt.Println("		EXP:" + questTime.Quest.QuestReward.RewardExp)
	}
	if len(questTime.Quest.QuestReward.RewardLocalQuality) != 0 {
		fmt.Println("		LocalQ:" + questTime.Quest.QuestReward.RewardLocalQuality)
	}
	if len(questTime.Quest.QuestReward.RewardLocalQualityAdditional) != 0 {
		fmt.Println("		LocalQAdd:" + questTime.Quest.QuestReward.RewardLocalQualityAdditional)
	}
	if len(questTime.Quest.QuestReward.RewardBy) != 0 {
		fmt.Println("		LocalQ by:" + questTime.Quest.QuestReward.RewardBy)
	}
	if len(questTime.Quest.QuestReward.RewardItem) != 0 {
		fmt.Println("		Item:" + questTime.Quest.QuestReward.RewardItem)
	}

	fmt.Println("---------")

	file, err := os.OpenFile("quests.json", os.O_RDWR, 0644)
	checkError(err)
	defer file.Close()

	f, err := ioutil.ReadAll(file)
	checkError(err)
	var alQs model.AllQuests
	err = json.Unmarshal(f, &alQs.Quests)
	checkError(err)

	alQs.Quests = append(alQs.Quests, &questTime)
	newAlQs, err := json.MarshalIndent(&alQs.Quests, "", " ")
	checkError(err)
	ioutil.WriteFile("quests.json", newAlQs, 0666)

}

func HandleJsonBig(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var receivedQuest model.QuestStructure
	err = json.Unmarshal(body, &receivedQuest)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)

	var questTime model.QuestTime
	questTime.Time = time.Now().Unix()
	questTime.Quest = receivedQuest

	fmt.Println("---------")

	fmt.Println("Time is " + string(strconv.FormatInt(questTime.Time, 10)))
	fmt.Println("	 Character:" + questTime.Quest.Character)
	fmt.Println("	 Content:" + questTime.Quest.Content)
	if len(questTime.Quest.QuestReward.QuestgiverName) != 0 {
		fmt.Println("	 	QG:" + questTime.Quest.QuestReward.QuestgiverName)
	}
	if len(questTime.Quest.QuestReward.RewardLp) != 0 {
		fmt.Println("		LP:" + questTime.Quest.QuestReward.RewardLp)
	}
	if len(questTime.Quest.QuestReward.RewardExp) != 0 {
		fmt.Println("		EXP:" + questTime.Quest.QuestReward.RewardExp)
	}
	if len(questTime.Quest.QuestReward.RewardLocalQuality) != 0 {
		fmt.Println("		LocalQ:" + questTime.Quest.QuestReward.RewardLocalQuality)
	}
	if len(questTime.Quest.QuestReward.RewardLocalQualityAdditional) != 0 {
		fmt.Println("		LocalQAdd:" + questTime.Quest.QuestReward.RewardLocalQualityAdditional)
	}
	if len(questTime.Quest.QuestReward.RewardBy) != 0 {
		fmt.Println("		LocalQ by:" + questTime.Quest.QuestReward.RewardBy)
	}
	if len(questTime.Quest.QuestReward.RewardItem) != 0 {
		fmt.Println("		Item:" + questTime.Quest.QuestReward.RewardItem)
	}

	fmt.Println("---------")

	file, err := os.OpenFile("quests.json", os.O_RDWR, 0644)
	checkError(err)
	defer file.Close()

	f, err := ioutil.ReadAll(file)
	checkError(err)
	//var alQs model.AllQuests
	var rootStr model.RootStructure
	err = json.Unmarshal(f, &rootStr.Root)
	//err = json.Unmarshal(f, &alQs.Quests)
	checkError(err)
	rootStr.Root.Quests = append(rootStr.Root.Quests, &questTime)

	//alQs.Quests = append(alQs.Quests, &questTime)
	newAlQs, err := json.MarshalIndent(&rootStr.Root, "", " ")
	checkError(err)
	ioutil.WriteFile("quests.json", newAlQs, 0666)

}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
