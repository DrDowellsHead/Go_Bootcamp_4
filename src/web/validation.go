package web

import (
	"encoding/json"
	"errors"
)

// Проверяю размер поля отдельно: при декодировании в массив Go
// недостающие элементы заполняются нулями, а лишние отбрасываются.
func validateBoardShape(body []byte) error {
	var input struct {
		Board json.RawMessage `json:"board"`
	}
	if err := json.Unmarshal(body, &input); err != nil {
		return err
	}
	var rows [][]json.RawMessage
	if len(input.Board) == 0 || string(input.Board) == "null" {
		return errors.New("board is required")
	}
	if err := json.Unmarshal(input.Board, &rows); err != nil {
		return errors.New("board must be a 3x3 integer matrix")
	}
	if len(rows) != 3 {
		return errors.New("board must have exactly 3 rows")
	}
	for _, row := range rows {
		if len(row) != 3 {
			return errors.New("each board row must have exactly 3 cells")
		}
	}
	return nil
}
