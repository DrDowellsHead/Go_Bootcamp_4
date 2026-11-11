package datasource

import "github.com/DrDowellsHead/Go_Bootcamp_4/src/domain"

func toRecord(game domain.Game) GameRecord {
	return GameRecord{ID: game.ID, Cells: [3][3]int(game.Board)}
}

func toDomain(record GameRecord) domain.Game {
	return domain.Game{ID: record.ID, Board: domain.Board(record.Cells)}
}
