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
	/*for _, val := range allQuests.Quests {
		fmt.Println(val.Quest.QuestReward)
	}*/
	//fmt.Println(allQuests)
	return &allQuests
}

//Todo: add `json:"time,omitempty"` to all models. there are empty fields in json now

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
	RewardLp                     string `json:"rewardLp,omitempty"`
	RewardExp                    string `json:"rewardExp,omitempty"`
	RewardLocalQuality           string `json:"rewardLocalQuality,omitempty"`
	RewardLocalQualityAdditional string `json:"rewardLocalQualityAdditional,omitempty"`
	RewardBy                     string `json:"rewardBy, omitempty"`
	RewardItem                   string `json:"rewardItem,omitempty"`
}
