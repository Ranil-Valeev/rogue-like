package generation

import (
	"rog/src/internal/game"
	"rog/src/internal/game/entities"
	"rog/src/internal/game/settings"
	"rog/src/internal/game/utils"
)

func createRooms() []*entities.Room {
	rooms := make([]*entities.Room, settings.RoomCount)
	sectorWidth := game.WorldWidth / settings.RoomGridSize
	sectorHeight := game.WorldHeight / settings.RoomGridSize

	marginX := 3
	marginY := 2

	for i := 0; i < settings.RoomGridSize; i++ {
		for j := 0; j < settings.RoomGridSize; j++ {
			sectorX1 := j*sectorWidth + 1
			sectorY1 := i*sectorHeight + 1
			sectorX2 := (j+1)*sectorWidth - 1
			sectorY2 := (i+1)*sectorHeight - 1

			roomWidth := utils.RandomFrom1ToN(4) + 10
			roomHeight := utils.RandomFrom1ToN(3) + 6

			maxXOffset := sectorWidth - roomWidth - 2*marginX
			if maxXOffset < 1 {
				maxXOffset = 1
			}

			maxYOffset := sectorHeight - roomHeight - 2*marginY
			if maxYOffset < 1 {
				maxYOffset = 1
			}

			x1 := sectorX1 + marginX + utils.RandomFrom1ToN(uint(maxXOffset))
			y1 := sectorY1 + marginY + utils.RandomFrom1ToN(uint(maxYOffset))
			x2 := x1 + roomWidth - 1
			y2 := y1 + roomHeight - 1

			if x2 >= sectorX2 {
				x2 = sectorX2 - 1
			}

			if y2 >= sectorY2 {
				y2 = sectorY2 - 1
			}

			rooms[i*settings.RoomGridSize+j] = entities.NewRoom(i*settings.RoomGridSize+j, x1, y1, x2, y2)
		}
	}

	return rooms
}

func carveRoom(w *entities.World, rooms []*entities.Room) {
	for _, r := range rooms {
		for x := r.GetX1(); x <= r.GetX2(); x++ {
			w.SetCell(x, r.GetY1(), game.CellRoomWall)
			w.SetCell(x, r.GetY2(), game.CellRoomWall)
		}

		for y := r.GetY1(); y <= r.GetY2(); y++ {
			w.SetCell(r.GetX1(), y, game.CellRoomWall)
			w.SetCell(r.GetX2(), y, game.CellRoomWall)
		}

		for y := r.GetY1() + 1; y < r.GetY2(); y++ {
			for x := r.GetX1() + 1; x < r.GetX2(); x++ {
				w.SetCell(x, y, game.CellFloor)
			}
		}
	}
}
