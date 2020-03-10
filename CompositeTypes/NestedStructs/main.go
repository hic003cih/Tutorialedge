package main

import "fmt"

func main() {
	type Person struct {
		name string
		age  int
	}

	// our Team struct
	type Team struct {
		name    string
		players [2]Person
	}

	// declaring an empty 'Team'
	var myTeam Team
	fmt.Println(myTeam)

	players := [...]Person{Person{name: "Forrest", age: 18}, Person{name: "Gordon"}}
	// declaring a team with players
	celtic := Team{name: "Celtic FC", players: players}
	fmt.Println(celtic)
}
