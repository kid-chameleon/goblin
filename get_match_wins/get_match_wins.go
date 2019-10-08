package main

import (
	"fmt"
	"os"

	"github.com/mchrobak/goblin"
)

func main() {
	shard := os.Getenv("PUBGSHARD")
	apiKey := os.Getenv("PUBGAPIKEY")
	pubgAPI, err := goblin.NewAPIClient(apiKey)
	if err != nil {
		panic(err.Error())
	}

	if _, err := pubgAPI.GetStatus(); err != nil {
		panic(err.Error())
	}
	fmt.Println("[PUBG PUBLIC API EXAMPLE]")

	players, err := pubgAPI.GetPlayers(os.Getenv("PUBGPLAYERNAME"), os.Getenv("PUBGPLAYERID"), shard)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Hello,", players.Data[0].Attributes.Name)
	playerID := players.Data[0].ID
	last := 0
	if len(players.Data[0].Relationships.Matches.Data) > 10 {
		last = 10
	} else {
		last = len(players.Data[0].Relationships.Matches.Data)
	}
	fmt.Printf("Checking for the number of wins in your last %d matches.\n", last)
	lastX := players.Data[0].Relationships.Matches.Data[:last]

	wonMatches := 0
	for _, match := range lastX {
		matchID := match.ID
		m, err := pubgAPI.GetMatch(matchID, shard)
		if err != nil {
			panic(err.Error())
		}

		for _, participant := range m.Included {
			if participant.Type == "participant" && participant.Attributes.Stats.PlayerID == playerID {
				if participant.Attributes.Stats.WinPlace == 1 {
					wonMatches++
				}
			}
		}
	}

	if wonMatches > 0 {
		fmt.Printf("You won %d of your last %d matches.\n", wonMatches, last)
	} else {
		fmt.Printf("You did not win any of your last %d games.\n", last)
	}
}
