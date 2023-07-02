package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetQuests() (aq *AllQuests) {
	res, err := http.Get("http://localhost:3000/api/getall")
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return &AllQuests{}
	}
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to parse response body")
		return &AllQuests{}
	}

	var allQuests AllQuests
	err = json.Unmarshal(body, &allQuests)
	if err != nil {
		fmt.Println("Failed to parse response json")
		return &AllQuests{}
	}
	return &allQuests
}

func GetQuestgiversQualities() []QuestgiverQualities {
	var questGiversQuelities []QuestgiverQualities

	res, err := http.Get("http://localhost:3000/api/questgiver/all")
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return questGiversQuelities
	}
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to parse response body")
		return questGiversQuelities
	}

	var receivedQuestgivers QuestgiversString
	err = json.Unmarshal(body, &receivedQuestgivers)
	if err != nil {
		fmt.Println("Failed to parse response json")
		return questGiversQuelities
	}
	fmt.Println(receivedQuestgivers)

	for _, el := range receivedQuestgivers.Questgivers {
		qg, err := http.Get("http://localhost:3000/api/questgiver/qualities/" + el)
		if err != nil {
			fmt.Printf("error making http request: %s\n on %s", err, el)
			return questGiversQuelities
		}
		body, err := ioutil.ReadAll(qg.Body)
		if err != nil {
			fmt.Println("Failed to parse response body of [" + el + "]")
			return questGiversQuelities
		}
		var receivedQuestgiverQualities QuestgiverQualities
		err = json.Unmarshal(body, &receivedQuestgiverQualities)
		if err != nil {
			fmt.Println("Failed to parse response json of [" + el + "]")
			return questGiversQuelities
		}
		questGiversQuelities = append(questGiversQuelities, receivedQuestgiverQualities)
	}

	return questGiversQuelities
}

func GetQestgiverQualitiesQuests(npcname string) QuestgiverQualitiesQuests {

	//get string of qualities in order to simple draw it.
	qg, err := http.Get("http://localhost:3000/api/questgiver/qualities/" + npcname)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return QuestgiverQualitiesQuests{}
	}
	body, err := ioutil.ReadAll(qg.Body)
	if err != nil {
		fmt.Println("Failed to parse response body of /qg/qualities/" + npcname)
		return QuestgiverQualitiesQuests{}
	}
	var receivedQuestgiverQualitiesQuests QuestgiverQualitiesQuests
	err = json.Unmarshal(body, &receivedQuestgiverQualitiesQuests)
	if err != nil {
		fmt.Println("Failed to parse response json of /qg/qualities")
		return QuestgiverQualitiesQuests{}
	}

	qs, err := http.Get("http://localhost:3000/api/questgiver/quests/" + npcname)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return QuestgiverQualitiesQuests{}
	}
	body, err = ioutil.ReadAll(qs.Body)
	if err != nil {
		fmt.Println("Failed to parse response body of /qg/qualities/" + npcname)
		return QuestgiverQualitiesQuests{}
	}
	var allQuests AllQuests
	err = json.Unmarshal(body, &allQuests)
	if err != nil {
		fmt.Println("Failed to parse response json")
		return QuestgiverQualitiesQuests{}
	}

	receivedQuestgiverQualitiesQuests.Quests = allQuests.Quests

	return receivedQuestgiverQualitiesQuests
}

// ********** unmarshalling models *************/
// /api/questgiver/questgivers
type QuestgiversString struct {
	Questgivers []string `json:"qgs"`
}

type Questgivers struct {
	Questgivers []QuestDescription `json:"ql"`
}

// /api/questgiver/generic
type QuestgiverQualities struct {
	Questgiver string   `json:"qg"`
	Qualities  []string `json:"ql"`
}

type QuestgiverQualitiesQuests struct {
	Questgiver string   `json:"qg"`
	Qualities  []string `json:"ql"`
	Quests     []*QuestTime
}

type QualitiesString struct {
	Qualities []string `json:"qls"`
}

// ******** base models ********/
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
