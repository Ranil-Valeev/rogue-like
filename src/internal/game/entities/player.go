package entities

import (
	"rog/src/internal/game"
)

type Player struct {
	x, y         int
	occupiedCell rune
	health       int
	maxHealth    int
	dexterity    int
	strength     int
	treasures    int
}

func NewPlayer(x, y int) *Player {
	return &Player{
		x:            x,
		y:            y,
		occupiedCell: game.CellFloor,
		health:       100,
		maxHealth:    100,
		dexterity:    100,
		strength:     100,
		treasures:    0,
	}
}

func (p *Player) GetX() int  { return p.x }
func (p *Player) SetX(x int) { p.x = x }

func (p *Player) GetY() int  { return p.y }
func (p *Player) SetY(y int) { p.y = y }

func (p *Player) GetPosition() (int, int) { return p.x, p.y }
func (p *Player) SetPosition(x, y int)    { p.x = x; p.y = y }

func (p *Player) GetHealth() int  { return p.health }
func (p *Player) SetHealth(x int) { p.health = x }

func (p *Player) GetMaxHealth() int  { return p.health }
func (p *Player) SetMaxHealth(x int) { p.health = x }

func (p *Player) GetOccupiedCell() rune     { return p.occupiedCell }
func (p *Player) SetOccupiedCell(cell rune) { p.occupiedCell = cell }

func (p *Player) GetDexterity() int  { return p.dexterity }
func (p *Player) SetDexterity(x int) { p.dexterity = x }

func (p *Player) GetStrenght() int  { return p.strength }
func (p *Player) SetStrenght(x int) { p.strength = x }

func (p *Player) GetTreasures() int  { return p.treasures }
func (p *Player) SetTreasures(x int) { p.treasures = x }
