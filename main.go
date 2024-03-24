package main

import (
    "fmt"
    "strconv"
)

func anyNil(slice []*string) bool {
    for _, elem := range slice {
        if elem == nil {
            return true
        }
    }
    return false
}

func noNil(slice []*string) bool {
    return !anyNil(slice)
}

func every(slice []*string) bool {
    // NOTE: we will never have a slice of length 0
    charToMatch := *slice[0]
    for _, elem := range slice {
        if elem == nil || *elem != charToMatch {
            return false
        }
    }
    return true
}

type Game struct {
    currentPlayer string
    board []*string
    turnsPlayed int
}

func (g *Game) Render () {
    fmt.Println()
    fmt.Printf("  %s's turn\n", g.currentPlayer)
    fmt.Println()
    renderString := ""
    for index, value := range g.board {
        var c string
        if value == nil {
            c = strconv.Itoa(index + 1)
        } else {
            c = *value
        }
        renderString = renderString + " " + c + " "
        if index != 2 && index != 5 && index != 8 {
            renderString = renderString + "|"
        }
        if index == 2 || index == 5 {
            fmt.Println(renderString)
            fmt.Println("-----------")
            renderString = ""
        }
    }
    fmt.Println(renderString)
    fmt.Println()
}

func (g *Game) findWinner() *string {
    regions := [][]*string{
        g.board[0:3], // top row
        g.board[3:6], // middle row
        g.board[6:9], // bottom row
        []*string{g.board[0], g.board[3], g.board[6]}, // left col
        []*string{g.board[1], g.board[4], g.board[7]}, // middle col
        []*string{g.board[2], g.board[5], g.board[8]}, // right col
        []*string{g.board[0], g.board[4], g.board[8]}, // top left to bottom right diagonal
        []*string{g.board[2], g.board[4], g.board[6]}, // top right to bottom left diagonal
    }
    for _, region := range regions {
        if noNil(region) && every(region) { return region[0] }
    }

    // no more moves and no winner is a tie
    if noNil(g.board) {
        tie := "Tie"
        return &tie
    }
    return nil
}

func NewGame() *Game {
    return &Game{
        currentPlayer: "x",
        board: make([]*string, 9),
        turnsPlayed: 0,
    }
}

func (g *Game) TakeTurn() {
    userInput := ""
    var cell int
    var err error
    for userInput == "" {
        fmt.Print("Enter the number of an unoccupied cell: ")
        _, err = fmt.Scanln(&userInput)
        if err != nil {
            fmt.Println("Error reading input:", err)
            return
        }
        // must be int
        cell, err = strconv.Atoi(userInput)
        if err != nil {
            fmt.Println("You must enter one of the remaining cell numbers")
            userInput = ""
            continue
        }
        // must be in range
        if cell < 1 || cell > 9 {
            fmt.Println("You must enter one of the remaining cell numbers")
            userInput = ""
            continue
        }
        //    must not be occupied
        if g.board[cell - 1] != nil {
            fmt.Println("That cell is already taken")
            userInput = ""
            continue
        }
        break
    }
    // update board
    playerValue := g.currentPlayer
    g.board[cell - 1] = &playerValue

    // swap player
    if g.currentPlayer == "x" {
        g.currentPlayer = "o"
    } else {
        g.currentPlayer = "x"
    }

    // render board
    g.Render()
}

func (g *Game) Play() {
    g.Render()
    winner := g.findWinner()
    for winner == nil {
        g.TakeTurn()
        winner = g.findWinner()
    }
    fmt.Println()
    fmt.Printf("The winner is %s!", *winner)
    fmt.Println()
}

func main() {
    g := NewGame()
    g.Play()
}
