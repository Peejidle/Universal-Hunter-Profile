// Package monster loads and defines monster data for the Universal Hunter Profile app.
package monster

import (
	"encoding/json"
	"fmt"
	"os"
)

type Monster struct {
	MonsterName string   `json:"name"`
	MonsterType string   `json:"type"`
	MonsterGame []string `json:"games"`
}

func LoadMonsters() ([]Monster, error) {
	rawData, err := os.ReadFile("monsterData.json")
	if err != nil {
		fmt.Println("Failed to read data:", err)
		return nil, err
	}

	var monsters []Monster
	err = json.Unmarshal(rawData, &monsters)
	if err != nil {
		fmt.Println("Failed to Unmarshal data:", err)
		return nil, err
	}
	return monsters, err
}
