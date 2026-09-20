// Package hunter defines what a hunter is and what it stores to be used elsewhere
package hunter

import "github.com/google/uuid"

type Hunter struct {
	HunterName   string
	HunterGame   string
	HunterID     uuid.UUID
	MonsterHunts []MonsterHunt
	WeaponUsages []WeaponUsage
}

type MonsterHunt struct {
	MonsterName   string
	MonsterAmount int
}

type WeaponUsage struct {
	WeaponName   string
	WeaponAmount int
}

func NewHunter(name, game string) Hunter {
	return Hunter{HunterName: name, HunterGame: game, HunterID: uuid.New(), MonsterHunts: []MonsterHunt{}, WeaponUsages: []WeaponUsage{}}
}

func RecordMonsterHunt(hunter Hunter, monsterName string) Hunter {
	found := false
	for i, hunt := range hunter.MonsterHunts {
		if monsterName == hunt.MonsterName {
			found = true
			hunter.MonsterHunts[i].MonsterAmount++
		}
	}
	if !found {
		hunter.MonsterHunts = append(hunter.MonsterHunts, MonsterHunt{MonsterName: monsterName, MonsterAmount: 1})
	}
	return hunter
}
