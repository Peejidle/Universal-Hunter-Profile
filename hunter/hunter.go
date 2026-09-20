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
