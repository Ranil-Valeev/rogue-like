package entities

import (
	"rog/src/internal/game"
	"rog/src/internal/game/settings"
)

type Monster struct {
	typeMonster  string
	x, y         int
	occupiedCell rune
	health       int
	maxHealth    int
	dexterity    int
	strength     int
	hostility    int
	roomId       int
	state        monsterState
}

type monsterState struct {
	moveDx  int
	moveDy  int
	visible bool
}

func NewMonster(x, y int, monsterType string, level int) *Monster {
	base, ok := settings.MonsterStats[monsterType]
	if !ok {
		base = settings.MonsterStats[settings.TypeZombie]
	}
	stats := settings.ScaleStatsForLevel(base, level)
	return &Monster{
		typeMonster:  monsterType,
		x:            x,
		y:            y,
		occupiedCell: game.CellFloor,
		health:       stats.HP,
		maxHealth:    stats.HP,
		dexterity:    stats.Dex,
		strength:     stats.Str,
		hostility:    stats.Hostility,
		state:        monsterState{visible: true},
	}
}

func (m *Monster) GetX() int { return m.x }
func (m *Monster) GetY() int { return m.y }

func (m *Monster) SetX(x int) { m.x = x }
func (m *Monster) SetY(y int) { m.y = y }

func (m *Monster) GetType() string {
	return m.typeMonster
}

func (m *Monster) GetPosition() (int, int) {
	return m.x, m.y
}
func (m *Monster) SetPosition(x, y int) {
	m.x = x
	m.y = y
}

func (m *Monster) GetOccupiedCell() rune {
	return m.occupiedCell
}

func (m *Monster) SetOccupiedCell(cell rune) {
	m.occupiedCell = cell
}

func (m *Monster) GetHealth() int  { return m.health }
func (m *Monster) SetHealth(x int) { m.health = x }

func (m *Monster) GetMaxHealth() int { return m.maxHealth }
func (m *Monster) GetDexterity() int { return m.dexterity }
func (m *Monster) GetStrength() int  { return m.strength }
func (m *Monster) GetHostility() int { return m.hostility }
func (m *Monster) IsAlive() bool     { return m.health > 0 }
func (m *Monster) GetRoomId() int    { return m.roomId }
func (m *Monster) SetRoomId(id int)  { m.roomId = id }

func (m *Monster) MoveDir() (int, int)     { return m.state.moveDx, m.state.moveDy }
func (m *Monster) SetMoveDir(dx, dy int)   { m.state.moveDx, m.state.moveDy = dx, dy }
func (m *Monster) IsVisible() bool         { return m.state.visible }
func (m *Monster) SetVisible(visible bool) { m.state.visible = visible }

func (m *Monster) MapSymbol() rune {
	if m.GetType() == settings.TypeGhost && !m.IsVisible() {
		return m.GetOccupiedCell()
	}
	if sym, ok := settings.MonsterSymbols[m.GetType()]; ok {
		return sym
	}
	return '?'
}
