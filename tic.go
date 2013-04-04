package main
import "fmt"
var b  = [9]string{"_","_","_","_","_","_","_","_","_"}
var t, i int = 1, 9
func main(){
	for(!check() && t < 10){
		t, i = (t+1), 9
		for(notIn(i)){
			fmt.Printf("%s %v %s", "Player ", (t%2+1), ", Please choose an unoccupied space (0-8): \n")
            fmt.Scan(&i)
		}
		if (t%2+1) == 1 { b[i] = "x" } else { b[i] = "o" }
		show()
	}
	if(check()) { fmt.Printf("%s %v %s", "Player ", (t%2+1), " wins.") } else { fmt.Printf("Tie.") }
}
func show(){
	fmt.Printf(b[0] + b[1] + b[2] + "\n" + b[3] + b[4] + b[5] + "\n" + b[6] + b[7] + b[8] + "\n")
}
func check() bool {
	if( (b[0] == b[1] && b[0] == b[2] && b[0] != "_") || (b[3] == b[4] && b[3] == b[5] && b[3] != "_") || 
		(b[6] == b[7] && b[6] == b[8] && b[6] != "_") || (b[0] == b[3] && b[0] == b[6] && b[0] != "_") || 
		(b[1] == b[4] && b[1] == b[7] && b[1] != "_") || (b[2] == b[5] && b[2] == b[8] && b[2] != "_") || 
		(b[0] == b[4] && b[0] == b[8] && b[0] != "_") || (b[2] == b[4] && b[2] == b[6] && b[2] != "_")) { return true }
	return false
}
func notIn(x int) bool {
	if ((x <= 8) && (x >= 0)) {	if ((b[x] == "_")) {return false} }
	return true
}