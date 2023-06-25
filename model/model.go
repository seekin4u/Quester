package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func ShowAllQuests() (aq *AllQuests) {
	file, err := os.OpenFile("quests.json", os.O_RDWR|os.O_APPEND, 0666)
	checkError(err)
	b, err := ioutil.ReadAll(file)
	var allQuests AllQuests
	json.Unmarshal(b, &allQuests.Quests)
	checkError(err)
	//fmt.Println(allQuests)
	return &allQuests
}

type AllQuests struct {
	Quests []*QuestTime
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
	RewardLp                     string `json:"rewardLp"`
	RewardExp                    string `json:"rewardExp"`
	RewardLocalQuality           string `json:"rewardLocalQuality"`
	RewardLocalQualityAdditional string `json:"rewardLocalQualityAdditional"`
	RewardItem                   string `json:"rewardItem"`
}
