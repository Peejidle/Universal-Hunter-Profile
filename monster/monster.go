// Package monster defines monster data and loading for the Universal Hunter Profile app.
package monster

type Monster struct {
	MonsterName string   `json:"name"`
	MonsterType string   `json:"type"`
	MonsterGame []string `json:"games"`
}

func loadMonster() {
}
