package entities

import (
	"rog/src/internal/game"
	"rog/src/internal/game/settings"
)

type World struct {
	width        int
	height       int
	cells        [][]rune
	currentLevel int
	Monsters     []*Monster
	Rooms        []*Room
}

func NewWorld() *World {
	w := &World{
		width:        game.WorldWidth,
		height:       game.WorldHeight,
		cells:        make([][]rune, game.WorldHeight),
		currentLevel: settings.DefaultStartLevel,
		Monsters:     make([]*Monster, 0),
	}

	for y := 0; y < game.WorldHeight; y++ {
		w.cells[y] = make([]rune, game.WorldWidth)
	}

	return w
}

func (w *World) GetCell(x, y int) rune {
	return w.cells[y][x]
}

func (w *World) SetCell(x, y int, symbol rune) {
	w.cells[y][x] = symbol
}

func (w *World) GetCurrentLevel() int {
	return w.currentLevel
}

func (w *World) SetCurrentLevel(n int) {
	w.currentLevel = n
}

func (w *World) AddMonster(m *Monster) {
	w.Monsters = append(w.Monsters, m)
}
func (w *World) GetRooms() []*Room      { return w.Rooms }
func (w *World) SetRooms(rooms []*Room) { w.Rooms = rooms }
func (w *World) GetRoomById(id int) *Room {
	for _, r := range w.Rooms {
		if r.GetId() == id {
			return r
		}
	}
	return nil
}
