package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/thedevsaddam/gojsonq/v2"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func ShowNpc(npcs *QuestStructure) {
	res := gojsonq.New().File("../data.json").From("items").Get()
	fmt.Printf("%#v\n", res)

}

func ShowAllQuests() (aq *AllQuests) {
	file, err := os.OpenFile("quests.json", os.O_RDWR|os.O_APPEND, 0666)
	checkError(err)
	b, err := ioutil.ReadAll(file)
	var allQuests AllQuests
	json.Unmarshal(b, &allQuests.Quests)
	checkError(err)
	return &allQuests
}

func ShowAllQuestsBig() (aq *RootStructure) {
	file, err := os.OpenFile("quests.json", os.O_RDWR|os.O_APPEND, 0666)
	checkError(err)
	b, err := ioutil.ReadAll(file)
	var rootSt RootStructure
	json.Unmarshal(b, &rootSt.Root)
	checkError(err)
	fmt.Println(rootSt)

	return &rootSt
}

func GetQgs() {

	res, err := http.Get("http://localhost:3000/api/qg")
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to parse response body")
		return
	}

	var receivedQuestgivers Questgivers
	err = json.Unmarshal(body, &receivedQuestgivers)
	if err != nil {
		fmt.Println("Failed to parse response json")
		return
	}
	fmt.Println(receivedQuestgivers)
}

func GetQestgiverQualities() QuestgiverQualities {

	res, err := http.Get("http://localhost:3000/api/questgiver/general")
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to parse response body")
		return QuestgiverQualities{}
	}

	var receivedQuestgiverQualities QuestgiverQualities
	err = json.Unmarshal(body, &receivedQuestgiverQualities)
	if err != nil {
		fmt.Println("Failed to parse response json")
		return QuestgiverQualities{}
	}
	fmt.Println(receivedQuestgiverQualities)
	return receivedQuestgiverQualities
}

// /api/questgiver/questgivers
type Questgivers struct {
	Questgivers []string `json:"qgs"`
}

// /api/questgiver/general
type QuestgiverQualities struct {
	Questgiver string   `json:"qg"`
	Qualities  []string `json:"ql"`
}

type RootStructure struct {
	Root AllQuests `json:"root"`
}

type AllQuests struct {
	Quests []*QuestTime `json:"array"`
}

type QuestTime struct {
	Time  int64          `json:"time"`
	Quest QuestStructure `json:"questStructure"`
}

type QuestStructure struct {
	Content     string           `json:"content"`
	Character   string           `json:"character"`
	QuestReward QuestDescription `json:"quest"`
}

type QuestDescription struct {
	QuestgiverName               string `json:"questgiverName"`
	RewardLp                     string `json:"rewardLp,omitempty"`
	RewardExp                    string `json:"rewardExp,omitempty"`
	RewardLocalQuality           string `json:"rewardLocalQuality,omitempty"`
	RewardLocalQualityAdditional string `json:"rewardLocalQualityAdditional,omitempty"`
	RewardBy                     string `json:"rewardBy,omitempty"`
	RewardItem                   string `json:"rewardItem,omitempty"`
}
