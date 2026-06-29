package monster

import "rog/src/internal/game/entities"

func ProcessTurn(w *entities.World, player *entities.Player) {
	for _, m := range w.Monsters {
		if !m.IsAlive() {
			continue
		}

		room := w.GetRoomById(m.GetRoomId())
		if room == nil {
			continue
		}

		moveByType(TurnContext{
			World:  w,
			Player: player,
			Room:   room,
		}, m)
	}
}
