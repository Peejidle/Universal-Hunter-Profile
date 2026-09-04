package monster

import (
	"encoding/json"
)


type Monster struct {
	MonsterName					string			`json:"name"`
	MonsterType					string			`json:"type"`
	MonsterGame					[]string		`json:"games"`
}


func loadMonster() {
	
}
