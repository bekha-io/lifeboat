package domain

import "log"

type Game struct {
	Players      []*Player
	Board        *Board
	CurrentPhase GamePhase
}

type GamePhase string

var (
	ProvisionDividingGamePhase = "Дележка припасов"
	ActionsGamePhase           = "Фаза действий"
	NavigationGamePhase        = "Фаза навигации"
)

type Player struct {
	Character      *Character
	Love           *LoveAndHateCard
	Hate           *LoveAndHateCard
	HandProvisions []Provision
}

func NewGame(players ...*Player) *Game {
	if len(players) > 8 || len(players) < 4 {
		log.Fatalln("There should be from 4 to 8 players to start a game")
	}

	var game *Game
	for _, player := range players {
		game.addPlayer(player)
	}

	return game
}

func (g *Game) addPlayer(player *Player) {
	g.Players = append(g.Players, player)
}
