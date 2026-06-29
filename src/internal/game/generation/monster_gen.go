package generation

import (
	"rog/src/internal/game"
	"rog/src/internal/game/entities"
	"rog/src/internal/game/settings"
	"rog/src/internal/game/utils"
)

var Config = settings.DefaultMonsterSpawnConfig()

func createMonster(room *entities.Room, monsterType string, level int, occupied map[int]bool) (*entities.Monster, bool) {
	x1 := room.GetX1() + 1
	y1 := room.GetY1() + 1
	x2 := room.GetX2() - 1
	y2 := room.GetY2() - 1

	width := x2 - x1 + 1
	height := y2 - y1 + 1

	for attempt := 0; attempt < width*height; attempt++ {
		x := x1 + utils.RandomFrom0ToN(width)
		y := y1 + utils.RandomFrom0ToN(height)
		key := x*10000 + y

		if occupied[key] {
			continue
		}

		occupied[key] = true
		monster := entities.NewMonster(x, y, monsterType, level)
		monster.SetRoomId(room.GetId())
		return monster, true
	}

	return nil, false
}

func getMonsterCount(level int) int {
	count := Config.BaseMonsters + (level*Config.MonsterPerLevel*3)/2
	extra := utils.RandomFrom0ToN(5) - 2
	count += extra

	if count < Config.MinMonsters {
		count = Config.MinMonsters
	}
	if count > Config.MaxMonsters {
		count = Config.MaxMonsters
	}

	return count
}

func getRandomTypeByPercent(level int) string {
	distribution := Config.GetConfigForLevel(level)

	var types []string
	for monsterType := range distribution {
		types = append(types, monsterType)
	}

	randNum := utils.RandomFrom0ToN(100)

	cumulative := 0
	for _, monsterType := range types {
		cumulative += distribution[monsterType]
		if randNum < cumulative {
			return monsterType
		}
	}

	return settings.TypeZombie
}

func CreateMonstersForLevel(rooms []*entities.Room, level int) []*entities.Monster {
	var monsters []*entities.Monster

	count := getMonsterCount(level)

	var availableRooms []*entities.Room
	for _, room := range rooms {
		if room.GetId() != game.Mid {
			availableRooms = append(availableRooms, room)
		}
	}

	if len(availableRooms) == 0 {
		return monsters
	}

	roomCounts := make(map[int]int)
	occupied := make(map[int]bool)

	spawned := 0
	for spawned < count && len(availableRooms) > 0 {
		idx := utils.RandomFrom0ToN(len(availableRooms))
		room := availableRooms[idx]

		monsterType := getRandomTypeByPercent(level)
		monster, ok := createMonster(room, monsterType, level, occupied)
		if !ok {
			availableRooms = append(availableRooms[:idx], availableRooms[idx+1:]...)
			continue
		}

		monsters = append(monsters, monster)
		spawned++

		roomCounts[room.GetId()]++
		if roomCounts[room.GetId()] >= Config.MaxMonstersPerRoom {
			availableRooms = append(availableRooms[:idx], availableRooms[idx+1:]...)
		}
	}

	return monsters
}

func carveMonsters(w *entities.World, monsters []*entities.Monster) {
	for _, m := range monsters {
		w.SetCell(m.GetX(), m.GetY(), m.MapSymbol())
	}
}
