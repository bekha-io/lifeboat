package domain

type Card interface {
	Title() string
	CardType() CardType
}

// CardType тип игровой карты (персонаж, припасы и т.д.)
type CardType string

var (
	CharacterCardType  CardType = "Персонаж"
	LoveCardType       CardType = "Лучший друг"
	HateCardType       CardType = "Заклятый враг"
	ProvisionCardType  CardType = "Провизия"
	WeatherCardType    CardType = "Погода"
	NavigationCardType CardType = "Навигация"
)
