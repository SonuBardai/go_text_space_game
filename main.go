package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Enemy struct {
	Name               string
	Attack             int
	Health             int
	EnemyDefeatMessage string
	EnemyWinMessage    string
}

type Room struct {
	Name        string
	Description string
	Enemies     []Enemy
	Exits       map[string]*Room
}

type Connection struct {
	From      string
	To        string
	Direction string
}

func oppositeDirection(direction string) string {
	directions := map[string]string{
		"north": "south",
		"south": "north",
		"east":  "west",
		"west":  "east",
	}
	return directions[direction]
}

func buildMap() *Room {
	rooms := []*Room{
		{
			Name:        "Cockpit",
			Description: "You find yourself in the cockpit of a spaceship. The controls are blinking and beeping, indicating that the ship is ready for takeoff.",
			Enemies:     make([]Enemy, 0),
			Exits:       make(map[string]*Room),
		},
		{
			Name:        "Cargo",
			Description: "You enter the ship’s cargo hold and find it filled with crates and containers. As you explore, you hear a strange noise coming from one of the crates.",
			Enemies:     make([]Enemy, 0),
			Exits:       map[string]*Room{},
		},
		{
			Name:        "Crate",
			Description: "You open the crate and find an alien creature inside! It has green skin, sharp teeth, and multiple tentacles.",
			Enemies:     []Enemy{{Name: "green tentacle monster", Health: 5, Attack: 8, EnemyDefeatMessage: "You manage to beat up the alien creature and lock it up in the cage it came out of. You can now move forward.", EnemyWinMessage: "The monster wraps you up with it's tentacles. You lose!"}},
			Exits:       map[string]*Room{},
		},
		{
			Name:        "Airlock",
			Description: "You find yourself in a dark, cramped airlock. The door to the spaceship is sealed shut, but you can see stars through the small window.",
			Enemies:     make([]Enemy, 0),
			Exits:       map[string]*Room{},
		},
		{
			Name:        "Asteroid",
			Description: "You manage to open the airlock door and float out into space. You see a nearby asteroid and decide to explore it.",
			Enemies:     make([]Enemy, 0),
			Exits:       map[string]*Room{},
		},
		{
			Name:        "EngineRoom",
			Description: "You find yourself in the engine room of a spaceship. The hum of the engines fills the air as you make your way through the maze of pipes and machinery.",
			Enemies:     make([]Enemy, 0),
			Exits:       map[string]*Room{},
		},
		{
			Name:        "ControlRoom",
			Description: "You enter the ship’s control room and find it deserted. You see that the navigator is broken, and needs some tools to repair it.",
			Enemies:     make([]Enemy, 0),
			Exits:       map[string]*Room{},
		},
		{
			Name:        "RepairRoom",
			Description: "The obvious place to find those tools are the repair room. You find the tools that are needed to fix the navigator.",
			Enemies:     make([]Enemy, 0),
			Exits:       map[string]*Room{},
		},
		{
			Name:        "Victory",
			Description: "You sit down at the controls and plot a course for home. The engines roar to life as the spaceship blasts off into hyperspace. Congratulations, you have won the game!",
			Enemies:     make([]Enemy, 0),
			Exits:       map[string]*Room{},
		},
		{
			Name:        "PlanetSurface",
			Description: "You step out of your spaceship and onto the surface of an alien planet. The air is thick and humid, and strange plants grow all around you.",
			Enemies:     make([]Enemy, 0),
			Exits:       map[string]*Room{},
		},
		{
			Name:        "Pirates",
			Description: "As you land on the asteroid, you are attacked by a group of space pirates! They are armed with laser guns and jetpacks, and they demand that you hand over your spaceship.",
			Enemies:     []Enemy{{Name: "space pirate", Attack: 20, Health: 14, EnemyDefeatMessage: "You manage to defeat the space pirate.", EnemyWinMessage: "The space pirates shoot youwitht heir lazers. You never had a chance. You lose!"}, {Name: "space pirate", Attack: 20, Health: 14, EnemyDefeatMessage: "You manage to defeat the space pirate.", EnemyWinMessage: "The space pirates shoot youwitht heir lazers. You never had a chance. You lose!"}, {Name: "space pirate", Attack: 20, Health: 14, EnemyDefeatMessage: "You manage to defeat the space pirate.", EnemyWinMessage: "The space pirates shoot youwitht heir lazers. You never had a chance. You lose!"}, {Name: "space pirate", Attack: 20, Health: 14, EnemyDefeatMessage: "You manage to defeat the space pirate.", EnemyWinMessage: "The space pirates shoot youwitht heir lazers. You never had a chance. You lose!"}, {Name: "space pirate", Attack: 20, Health: 14, EnemyDefeatMessage: "You manage to defeat the space pirate.", EnemyWinMessage: "The space pirates shoot youwitht heir lazers. You never had a chance. You lose!"}},
			Exits:       map[string]*Room{},
		},
		{
			Name:        "MedicalBay",
			Description: "You enter the medical bay of a spaceship.",
			Enemies:     make([]Enemy, 0),
			Exits:       map[string]*Room{},
		},
		{
			Name:        "Lab",
			Description: "You come across a laboratory filled with strange experiments and alien specimens.",
			Enemies:     make([]Enemy, 0),
			Exits:       map[string]*Room{},
		},
		{
			Name:        "Specimen",
			Description: "As you examine one of the specimens, it suddenly comes to life! It is a large, slimy creature with multiple eyes and tentacles.",
			Enemies:     []Enemy{{Name: "slimy creature", Attack: 4, Health: 10, EnemyDefeatMessage: "You slam the slimy creature and flush it down. You can now move on.", EnemyWinMessage: "You got swallowed by the creature. You lose."}},
			Exits:       map[string]*Room{},
		},
	}

	connections := []*Connection{
		{From: "Cockpit", To: "Cargo", Direction: "north"},
		{From: "Cargo", To: "Crate", Direction: "north"},
		{From: "Crate", To: "PlanetSurface", Direction: "east"},

		{From: "Crate", To: "EngineRoom", Direction: "north"},

		{From: "Cargo", To: "Airlock", Direction: "west"},
		{From: "Airlock", To: "Asteroid", Direction: "all"},
		{From: "Asteroid", To: "Pirates", Direction: "all"},

		{From: "Cargo", To: "EngineRoom", Direction: "east"},
		{From: "EngineRoom", To: "RepairRoom", Direction: "all"},
		{From: "RepairRoom", To: "Victory", Direction: "all"},

		{From: "Crate", To: "MedicalBay", Direction: "west"},
		{From: "MedicalBay", To: "Lab", Direction: "all"},
		{From: "Lab", To: "Specimen", Direction: "all"},
		{From: "Specimen", To: "Airlock", Direction: "north"},
		{From: "Specimen", To: "EngineRoom", Direction: "west"},
		{From: "Specimen", To: "EngineRoom", Direction: "east"},
	}

	for _, connection := range connections {
		var fromRoom, toRoom *Room
		for _, room := range rooms {
			if room.Name == connection.From {
				fromRoom = room
			} else if room.Name == connection.To {
				toRoom = room
			}

			if fromRoom != nil && toRoom != nil {
				break
			}
		}
		if connection.Direction == "all" {
			fromRoom.Exits["north"] = toRoom
			fromRoom.Exits["south"] = toRoom
			fromRoom.Exits["east"] = toRoom
			fromRoom.Exits["west"] = toRoom
		} else {
			fromRoom.Exits[connection.Direction] = toRoom
			toRoom.Exits[oppositeDirection(connection.Direction)] = fromRoom
		}
	}

	startRoom := rooms[0]
	return startRoom
}

type Player struct {
	CurrentRoom *Room
	Health      int
	Attack      int
}

func (player *Player) defeatedAllEnemies() bool {
	for _, enemy := range player.CurrentRoom.Enemies {
		if enemy.Health > 0 {
			return false
		}
	}
	return true
}

func (player *Player) printFightStats() {
	fmt.Printf("Your stats: Health (%d)  Attack (%d)\n", player.Health, player.Attack)
	for _, enemy := range player.CurrentRoom.Enemies {
		fmt.Printf("%s Stats: Health (%d)  Attack (%d)\n", enemy.Name, enemy.Health, enemy.Attack)
	}
}

func (player *Player) Move(direction string) {
	fmt.Print("\n")
	if len(player.CurrentRoom.Enemies) != 0 && !player.defeatedAllEnemies() {
		fmt.Println("You can't run away! You must fight!")
		player.printFightStats()
		return
	}
	if nextRoom, ok := player.CurrentRoom.Exits[direction]; ok {
		player.CurrentRoom = nextRoom
		fmt.Println(player.CurrentRoom.Name)
		fmt.Println(player.CurrentRoom.Description)

		if len(nextRoom.Enemies) != 0 && !player.defeatedAllEnemies() {
			player.printFightStats()
		}
	} else {
		fmt.Println("You can't go there!")
	}
}

func (player *Player) Fight(enemy string) {
	var enemyToAttack *Enemy
	for i, e := range player.CurrentRoom.Enemies {
		if e.Name == enemy {
			enemyToAttack = &player.CurrentRoom.Enemies[i]
			break
		}
	}

	if enemyToAttack == nil {
		for i, e := range player.CurrentRoom.Enemies {
			if e.Health > 0 {
				enemyToAttack = &player.CurrentRoom.Enemies[i]
			}
		}
	}

	enemyToAttack.Health -= player.Attack
	player.Health -= enemyToAttack.Attack
	if player.Health <= 0 {
		player.Health = 0
		fmt.Println(enemyToAttack.EnemyWinMessage)
		fmt.Println("GAME OVER!")
	}
	if enemyToAttack.Health <= 0 {
		enemyToAttack.Health = 0
		fmt.Println(enemyToAttack.EnemyDefeatMessage)
		if !player.defeatedAllEnemies() {
			player.printFightStats()
		}
	} else if player.Health > 0 {
		fmt.Printf("You attacked the %s. You caused %d damage. You took %d damage from the %s\n", enemyToAttack.Name, player.Attack, enemyToAttack.Attack, enemyToAttack.Name)
		player.printFightStats()
	}

}

func main() {
	startingRoom := buildMap()
	player := &Player{CurrentRoom: startingRoom, Health: 100, Attack: 10}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(player.CurrentRoom.Description)

loop:
	for scanner.Scan() {
		input := scanner.Text()
		command := strings.Fields(input)

		switch command[0] {
		case "quit":
			break loop
		case "exit":
			break loop
		case "go":
			if len(command) < 2 {
				fmt.Println("Go where?")
			} else {
				player.Move(command[1])
			}
		case "attack":
			if len(command) < 2 {
				fmt.Println("Attack who?")
			} else {
				player.Fight(command[1])
				if player.Health == 0 {
					break loop
				}
			}
		default:
			fmt.Println("You can't do that!")
		}
	}
}
