package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("OOXX Game")
	w.Resize(fyne.NewSize(300, 300))

	var currentPlayer = "X" // 默认从玩家 X 开始
	board := make([][]*widget.Button, 3)
	for i := range board {
		board[i] = make([]*widget.Button, 3)
	}

	grid := container.NewGridWithColumns(3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			btn := widget.NewButton("", nil)
			btn.OnTapped = makeButtonHandler(btn, board, &currentPlayer, a)
			board[i][j] = btn
			grid.Add(btn)
		}
	}

	w.SetContent(grid)
	w.ShowAndRun()
}

func makeButtonHandler(btn *widget.Button, board [][]*widget.Button, currentPlayer *string, a fyne.App) func() {
	return func() {
		if btn.Text == "" {
			btn.SetText(*currentPlayer)
			if checkWin(board, *currentPlayer) {
				a.SendNotification(fyne.NewNotification("Winner", "Player "+*currentPlayer+" wins!"))
				resetBoard(board)
			} else if isBoardFull(board) {
				a.SendNotification(fyne.NewNotification("Draw", "The game is a draw!"))
				resetBoard(board)
			}
			*currentPlayer = switchPlayer(*currentPlayer)
		}
	}
}

func checkWin(board [][]*widget.Button, player string) bool {
	for i := 0; i < 3; i++ {
		if (board[i][0].Text == player && board[i][1].Text == player && board[i][2].Text == player) ||
			(board[0][i].Text == player && board[1][i].Text == player && board[2][i].Text == player) {
			return true
		}
	}
	if (board[0][0].Text == player && board[1][1].Text == player && board[2][2].Text == player) ||
		(board[0][2].Text == player && board[1][1].Text == player && board[2][0].Text == player) {
		return true
	}
	return false
}

func isBoardFull(board [][]*widget.Button) bool {
	for _, row := range board {
		for _, btn := range row {
			if btn.Text == "" {
				return false
			}
		}
	}
	return true
}

func resetBoard(board [][]*widget.Button) {
	for _, row := range board {
		for _, btn := range row {
			btn.SetText("")
		}
	}
}

func switchPlayer(current string) string {
	if current == "X" {
		return "O"
	}
	return "X"
}
