package main

import (
	"fmt"
	term "github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

func reset() {
	term.Sync() // prettier layout
}

func printPattern(newPattern [6][8]string) {
	for i := 0; i < 6; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("%s", newPattern[i][j])
		}
		fmt.Printf("\n")
	}
}

func main() {
	// layouting pattern
	locArray := [6][8]string{
		{"#", "#", "#", "#", "#", "#", "#", "#"},
		{"#", ".", ".", ".", ".", ".", ".", "#"},
		{"#", ".", "#", "#", "#", ".", ".", "#"},
		{"#", ".", ".", ".", "#", ".", "#", "#"},
		{"#", ".", "#", ".", ".", ".", ".", "#"},
		{"#", "#", "#", "#", "#", "#", "#", "#"},
	}

	// create all probable treasure position
	treasureProbablePos := [17][2]int{
		{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6},
		{2, 1}, {2, 5}, {2, 6},
		{3, 1}, {3, 2}, {3, 3}, {3, 5},
		{4, 3}, {4, 4}, {4, 5}, {4, 6},
	}

	// random treasure position
	rand.Seed(time.Now().UnixNano())
	treasurePos := rand.Intn(16)                       // get random index of treasure Probable position
	rowTreasure := treasureProbablePos[treasurePos][0] // get the row pos
	colTreasure := treasureProbablePos[treasurePos][1] // get the col pos
	// end of random treasure position

	locArray[rowTreasure][colTreasure] = "$" // set the treasure position

	rowMin := 1    // minimal row pos
	rowMax := 4    // maximal row pos
	columnMax := 6 // maximal column pos

	newPattern := locArray

	// player start position
	row := 4
	column := 1
	newPattern[row][column] = "X" // set the player position

	err := term.Init()
	if err != nil {
		panic(err)
	}
	printPattern(newPattern) // print the initial pattern
	fmt.Println("Enter Up, Right, Down key or press ESC button to quit")
	lastMove := "start" // save the last move

keyPressListenerLoop:
	for {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEsc:
				break keyPressListenerLoop
			case term.KeyArrowUp: // if player press up key button
				if lastMove == "start" || lastMove == "up" { // always check move after pressed
					newPattern := locArray // copy the pattern
					if row-1 >= rowMin {   // check if the position is not override
						// move up, so the row index must be decrease by 1
						row -= 1
						if newPattern[row][column] != "#" {
							newPattern[row][column] = "X" // if the new position is not "#", set X in new position
							printPattern(newPattern)

							// if new position is the location of the treasure. game end and player win
							if newPattern[rowTreasure][colTreasure] == "X" {
								fmt.Println("YOU WIN!")
								break keyPressListenerLoop
							}
							reset()
							fmt.Println("Move up")
							lastMove = "up"
						} else {
							row += 1
							newPattern[row][column] = "X"
							printPattern(newPattern)

							fmt.Println("You can't move to that direction")
						}
					} else {
						fmt.Println("You can't move to that direction")
					}
				} else {
					fmt.Println("You can't move to up direction")
				}

			case term.KeyArrowRight:
				// move up is mandatory if you want to move right. so you have to move up first before move right
				// but  you can move right more than 1 times
				if lastMove == "up" || lastMove == "right" {
					newPattern := locArray
					if column+1 <= columnMax {
						// move right, so the column index have to be increased by 1
						column += 1
						if newPattern[row][column] != "#" { // if the new position is not "#", set X in new position
							newPattern[row][column] = "X"
							printPattern(newPattern)

							// if new position is the location of the treasure. game end and player win
							if newPattern[rowTreasure][colTreasure] == "X" {
								fmt.Println("YOU WIN!")
								break keyPressListenerLoop
							}
							reset()
							fmt.Println("Move Right")
							lastMove = "right"
						} else {
							column -= 1
							newPattern[row][column] = "X"
							printPattern(newPattern)

							fmt.Println("You can't move to that direction")
						}
					} else {
						fmt.Println("You can't move to that direction")
					}
				} else {
					fmt.Println("You can't move to right direction")
				}

			case term.KeyArrowDown:
				// move right is mandatory if you want to move right. so you have to move right first before move down
				// but  you can move down more than 1 times
				if lastMove == "right" || lastMove == "down" {
					newPattern := locArray
					if row+1 <= rowMax {
						// move dpwm, so the row index have to be increased by 1
						row += 1
						if newPattern[row][column] != "#" { // if the new position is not "#", set X in new position
							newPattern[row][column] = "X"
							printPattern(newPattern)

							// if new position is the location of the treasure. game end and player win
							if newPattern[rowTreasure][colTreasure] == "X" {
								fmt.Println("YOU WIN!")
								break keyPressListenerLoop
							}
							reset()
							fmt.Println("Move Down")
							lastMove = "down"
						} else {
							row -= 1
							newPattern[row][column] = "X"
							printPattern(newPattern)
							if newPattern[rowTreasure][colTreasure] == "X" {
								fmt.Println("YOU WIN!")
							}
							fmt.Println("You can't move to that direction")
						}
					} else {
						fmt.Println("You can't move to that direction")
					}
				} else {
					fmt.Println("You can't move to down direction")
				}

			default:
				// we only want to read a single character or one key pressed event
				reset()
				fmt.Println("Please just Up, Right, Down arrow key")

			}
		case term.EventError:
			panic(ev.Err)
		}
	}
}
