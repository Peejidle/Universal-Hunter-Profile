package monster

import "testing"

func TestLoadMonsters(t *testing.T) {
	monsters, err := LoadMonsters()
	if err != nil {
		t.Fatalf("LoadMonsters returned an error: %v", err)
	}

	if len(monsters) == 0 {
		t.Fatal("LoadMonsters returned an empty slice, expected monster data")
	}
}
