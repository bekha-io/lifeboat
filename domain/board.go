package domain

type Board struct {
	Boat           *Boat
	ProvisionDeck  []Provision
	WeatherDeck    []Weather
	NavigationDeck []Navigation
}

func (b *Board) CharactersPlayingNum() int {
	return len(b.Boat.Seats)
}

func NewBoard(playersNum int) *Board {
	var board *Board
	characters := getCharactersRandom(playersNum)
	board.Boat = NewBoat(characters...)
	board.ProvisionDeck = GetProvisionAll()
	board.WeatherDeck = GetWeatherAll()
	board.NavigationDeck = GetNavigationAll()
	return board
}
