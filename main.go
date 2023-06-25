package main

import (
	"log"
	"net/http"
	"os"

	handlers "Quester/handlers"
)

// TODO: Handlefunc for static html, save json to local in handleJson, display saved json objects in front html

func main() {

	//handlers.PrintAllQuests()
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

type AllQuests struct {
	Time   int64       `json:"time"`
	Quests []QuestJson `json:"allQuests"`
}

type QuestJson struct {
	Content     string `json:"content"`
	Character   string `json:"character"`
	QuestReward Quest  `json:"quest"`
}

type Quest struct {
	QuestgiverName               string `json:"questgiverName"`
	RewardLp                     string `json:"rewardLp"`
	RewardExp                    string `json:"rewardExp"`
	RewardLocalQuality           string `json:"rewardLocalQuality"`
	RewardLocalQualityAdditional string `json:"rewardLocalQualityAdditional"`
	RewardItem                   string `json:"rewardItem"`
}

//{[]}
