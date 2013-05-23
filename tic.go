package main

import "fmt"

// game board
var b  = [9]string{"_","_","_","_","_","_","_","_","_"}

// turn counter and selected move
var t, i int = 1, 9

func main(){
	// main game loop
	// game ends when the all spaces or filled or when a player wins
	for(!check() && t < 10){
		// increment turn number
		// and set move to an invalid value
		t, i = (t+1), 9
		
		// while move is invalid, ask for a valid move
		for(notIn(i)){
			fmt.Printf("%s %v %s", "Player ", (t%2+1), ", Please choose an unoccupied space (0-8): \n")
           		fmt.Scan(&i)
		}
		
		// if it is player 1's turn, write an 'x' to the board
		// otherwise, write an 'o' to the board
		if (t%2+1) == 1 { b[i] = "x" } else { b[i] = "o" }
		
		// show game board after move is taken
		show()
	}
	
	// if a player has won, display victory message
	// otherwise indicate a tie game
	if(check()) { fmt.Printf("%s %v %s", "Player ", (t%2+1), " wins.") } else { fmt.Printf("Tie.") }
}

// output current game board to screen
func show(){
	fmt.Printf(b[0] + b[1] + b[2] + "\n" + b[3] + b[4] + b[5] + "\n" + b[6] + b[7] + b[8] + "\n")
}

// check to see if a player has won
func check() bool {
	if( 	(b[0] == b[1] && b[0] == b[2] && b[0] != "_") || (b[3] == b[4] && b[3] == b[5] && b[3] != "_") || 
		(b[6] == b[7] && b[6] == b[8] && b[6] != "_") || (b[0] == b[3] && b[0] == b[6] && b[0] != "_") || 
		(b[1] == b[4] && b[1] == b[7] && b[1] != "_") || (b[2] == b[5] && b[2] == b[8] && b[2] != "_") || 
		(b[0] == b[4] && b[0] == b[8] && b[0] != "_") || (b[2] == b[4] && b[2] == b[6] && b[2] != "_")){ return true }
	return false
}

// check to see if the space selected for a player's move is already taken
// or if move is out of bounds
func notIn(x int) bool {
	if ((x <= 8) && (x >= 0)) { if ((b[x] == "_")) {return false} }
	return true
}
