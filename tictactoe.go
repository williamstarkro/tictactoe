package main

import (
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"os"
)

// 3 x 3

type line [3]string

type tictactoe [8]line

func startGame() *tictactoe {

	line := [...]string{"-","-","-"}
	ttt := &tictactoe{line,line,line,
	line,line,line,
	line,line}
	return ttt
}

func (l *line) checkForWinner() bool {
	if l[0] != "-" {
		if l[0] == l[1] && l[1] == l[2] {
			return true
		}
	}
	return false
}

func (ttt *tictactoe) checkAll() bool {
	for _, line := range ttt {
		if line.checkForWinner() {
			return true
		}
	}
	return false
}

func (ttt *tictactoe) acceptInput(x, y int, val string) bool {

	if ttt[x][y] != "-" {
		fmt.Println("This spot is taken")
		return false
	}

	ttt[x][y] = val
	ttt[x+3][y] = val
	if x == y  && x == 1 {
		ttt[6][1] = val
		ttt[7][1] = val
	} else if x == y {
		ttt[6][x] = val
	} else if x%2 == 0 && y%2 == 0 {
		ttt[7][y] = val
	}
	return true
}

func (ttt *tictactoe) print() {

	fmt.Printf("%s|%s|%s", ttt[0][0], ttt[0][1], ttt[0][2])
	fmt.Println()
	fmt.Printf("%s|%s|%s", ttt[1][0], ttt[1][1], ttt[1][2])
	fmt.Println()
	fmt.Printf("%s|%s|%s", ttt[2][0], ttt[2][1], ttt[2][2])
	fmt.Println()

}

func main() {

	ttt := startGame()
	scanner := bufio.NewScanner(os.Stdin)
	players := [2]string{"X", "O"}
	var player int = 0

	ttt.print()

	for {
		fmt.Printf("%s's turn: ", players[player])
		scanner.Scan()
		s := strings.Split(scanner.Text(), " ")
		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])
		if ttt.acceptInput(x,y,players[player]) {
			if ttt.checkAll() {
				fmt.Printf("%s is the winner", players[player])
				fmt.Println()
				ttt.print()
				return
			}
			player = (player + 1)%2
		}
		ttt.print()
	}
}















