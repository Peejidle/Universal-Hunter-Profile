package main

import (
	"fmt"

	"uhp/monster"
)

func main() {
	monsters, err := monster.LoadMonsters()
	if err != nil {
		fmt.Println("Failed to Load Monsters", err)
		return
	}
	for _, m := range monsters {
		fmt.Println(m.MonsterName)
	}
}
