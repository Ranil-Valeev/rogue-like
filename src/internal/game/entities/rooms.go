package entities

type Room struct {
	id            int
	x1, y1        int // левый верхний угол
	x2, y2        int // нижний правый угол
	doors         []Door
	amountOfDoors int
}

func NewRoom(id, x1, y1, x2, y2 int) *Room {
	return &Room{
		id:    id,
		x1:    x1,
		y1:    y1,
		x2:    x2,
		y2:    y2,
		doors: []Door{},
	}
}

func (r *Room) GetId() int {
	return r.id
}

func (r *Room) GetCenterX() int {
	return (r.x1 + r.x2) / 2
}

func (r *Room) GetCenterY() int {
	return (r.y1 + r.y2) / 2
}

func (r *Room) GetX1() int { return r.x1 }
func (r *Room) GetY1() int { return r.y1 }
func (r *Room) GetX2() int { return r.x2 }
func (r *Room) GetY2() int { return r.y2 }

func (r *Room) ContainsFloor(x, y int) bool {
	return x > r.x1 && x < r.x2 &&
		y > r.y1 && y < r.y2
}

func (r *Room) SetAmountOfDoors(n int) {
	r.amountOfDoors = n
}

func (r *Room) GetAmountOfDoors() int {
	return r.amountOfDoors
}

func (r *Room) SetDoor(x, y int) {
	r.doors = append(r.doors, Door{x: x, y: y})
}

func (r *Room) GetDoors() []Door {
	return r.doors
}

type Door struct {
	x, y      int
	connected bool
}

func (d *Door) GetX() int {
	return d.x
}

func (d *Door) GetY() int {
	return d.y
}

func (d *Door) SetConnected(x bool) {
	d.connected = x
}

func (d *Door) GetConnected() bool {
	return d.connected
}
