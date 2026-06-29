package generation

import (
	"rog/src/internal/game"
	"rog/src/internal/game/entities"
	"rog/src/internal/game/utils"
)

func createDoors(w *entities.World, rooms []*entities.Room) {
	for roomNum, room := range rooms {
		switch roomNum {
		case game.TopLeft, game.TopRight, game.BotLeft, game.BotRight:
			room.SetAmountOfDoors(1)
			carveDoors(roomNum, w, room)
		case game.Top, game.MidLeft, game.MidRight, game.Bot:
			room.SetAmountOfDoors(1)
			carveDoors(roomNum, w, room)
		case game.Mid:
			room.SetAmountOfDoors(1)
			carveDoors(roomNum, w, room)
		}
	}
}

func carveDoors(roomNum int, w *entities.World, room *entities.Room) {
	amountOfDoors := room.GetAmountOfDoors()
	availableDirs := getAvailableDirections(roomNum)

	for i := 0; i < amountOfDoors; i++ {
		idx := utils.RandomFrom1ToN(uint(len(availableDirs))) - 1
		dir := availableDirs[idx]

		availableDirs = append(availableDirs[:idx], availableDirs[idx+1:]...)

		x, y := findDoorPosition(room, dir)

		room.SetDoor(x, y)
		w.SetCell(x, y, game.CellRoomDoor)
	}
}

func getAvailableDirections(roomNum int) []int {
	switch roomNum {
	case game.TopLeft:
		return []int{game.Right, game.Down}
	case game.TopRight:
		return []int{game.Left, game.Down}
	case game.BotLeft:
		return []int{game.Right, game.Up}
	case game.BotRight:
		return []int{game.Left, game.Up}

	case game.Top:
		return []int{game.Left, game.Right, game.Down}
	case game.Bot:
		return []int{game.Left, game.Right, game.Up}
	case game.MidLeft:
		return []int{game.Up, game.Down, game.Right}
	case game.MidRight:
		return []int{game.Up, game.Down, game.Left}

	case game.Mid:
		return []int{game.Up, game.Right, game.Down, game.Left}
	}

	return []int{-1, -1, -1, -1}
}

func findDoorPosition(room *entities.Room, dir int) (int, int) {
	centerX := room.GetCenterX()
	centerY := room.GetCenterY()
	x1, y1 := room.GetX1(), room.GetY1()
	x2, y2 := room.GetX2(), room.GetY2()

	switch dir {
	case game.Up:
		return centerX, y1

	case game.Down:
		return centerX, y2

	case game.Left:
		return x1, centerY

	case game.Right:
		return x2, centerY
	}

	return -1, -1
}
