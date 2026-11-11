package web

import "github.com/DrDowellsHead/Go_Bootcamp_4/src/domain"

func toDomain(game Game) domain.Game {
	return domain.Game{ID: game.ID, Board: domain.Board(game.Board)}
}

func toWeb(game domain.Game, service domain.Service) Game {
	finished, winner := service.IsFinished(game)
	status := "playing"
	if finished {
		if winner == domain.Empty {
			status = "draw"
		} else {
			status = "finished"
		}
	}
	return Game{ID: game.ID, Board: [3][3]int(game.Board), Status: status, Winner: winner}
}
