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
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func PrintAllQuests() {
	file, err := os.OpenFile("quests.json", os.O_RDWR|os.O_APPEND, 0666)
	checkError(err)
	b, err := ioutil.ReadAll(file)
	var allQuests model.AllQuests
	json.Unmarshal(b, &allQuests.Quests)

	for _, v := range allQuests.Quests {

		fmt.Println("---------")

		fmt.Println("Time is " + string(strconv.FormatInt(v.Time, 10)))
		fmt.Println("	 Character:" + v.Quest.Character)
		fmt.Println("	 Content:" + v.Quest.Content)
		if len(v.Quest.QuestReward.QuestgiverName) != 0 {
			fmt.Println("	 	QG:" + v.Quest.QuestReward.QuestgiverName)
		}
		if len(v.Quest.QuestReward.RewardLp) != 0 {
			fmt.Println("		LP:" + v.Quest.QuestReward.RewardLp)
		}
		if len(v.Quest.QuestReward.RewardExp) != 0 {
			fmt.Println("		EXP:" + v.Quest.QuestReward.RewardExp)
		}
		if len(v.Quest.QuestReward.RewardLocalQuality) != 0 {
			fmt.Println("		LocalQ:" + v.Quest.QuestReward.RewardLocalQuality)
		}
		if len(v.Quest.QuestReward.RewardLocalQualityAdditional) != 0 {
			fmt.Println("		LocalQAdd:" + v.Quest.QuestReward.RewardLocalQualityAdditional)
		}
		if len(v.Quest.QuestReward.RewardItem) != 0 {
			fmt.Println("		Item:") //TODO
		}

		fmt.Println("---------")
	}
	fmt.Println(cap(allQuests.Quests))
	fmt.Println(len(allQuests.Quests))

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	au := model.ShowAllQuests()
	//t, err := template.ParseFiles("templates/index.html")
	t := template.Must(template.ParseFiles("templates/index.html"))
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

	fmt.Println("--------- aquired json")

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
	if len(questTime.Quest.QuestReward.RewardItem) != 0 {
		//fmt.Println("		Item:") //TODO
	}

	fmt.Println("---------")

	file, err := os.OpenFile("quests.json", os.O_RDWR, 0644)
	checkError(err)
	defer file.Close()

	f, err := ioutil.ReadAll(file)
	var alQs model.AllQuests
	err = json.Unmarshal(f, &alQs.Quests)
	checkError(err)

	alQs.Quests = append(alQs.Quests, &questTime)
	newAlQs, err := json.MarshalIndent(&alQs.Quests, "", " ")
	checkError(err)
	ioutil.WriteFile("quests.json", newAlQs, 0666)

}
