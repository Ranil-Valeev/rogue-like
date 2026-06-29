package monster

import "rog/src/internal/game/entities"

type TurnContext struct {
	World  *entities.World
	Player *entities.Player
	Room   *entities.Room
}
