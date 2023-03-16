package domain

import "fmt"

type LoveAndHateCard struct {
	IsLove    bool
	Character Character
}

func (l LoveAndHateCard) IsHate() bool {
	return !l.IsLove
}

func (l LoveAndHateCard) CardType() CardType {
	if l.IsLove {
		return LoveCardType
	}
	return HateCardType
}

func (l LoveAndHateCard) Title() string {
	if l.IsHate() {
		return fmt.Sprintf("Заклятый враг (%v)", l.Character.Name)
	}
	return fmt.Sprintf("Лучший друг (%v)", l.Character.Name)
}

// Value количество очков, которое достанется владельцу карты при выживании / смерти указанного на карте персонажа
func (l LoveAndHateCard) Value() int {
	return l.Character.SurvivalValue
}

var LoveTheKid = LoveAndHateCard{
	IsLove:    true,
	Character: TheKid,
}

var HateTheKid = LoveAndHateCard{
	IsLove:    false,
	Character: TheKid,
}

var LoveLadyLauren = LoveAndHateCard{
	IsLove:    true,
	Character: LadyLauren,
}

var HateLadyLauren = LoveAndHateCard{
	IsLove:    false,
	Character: LadyLauren,
}

var LoveSirStephen = LoveAndHateCard{
	IsLove:    true,
	Character: SirStephen,
}

var HateSirStephen = LoveAndHateCard{
	IsLove:    false,
	Character: SirStephen,
}

var LoveTheCaptain = LoveAndHateCard{
	IsLove:    true,
	Character: TheCaptain,
}

var HateTheCaptain = LoveAndHateCard{
	IsLove:    false,
	Character: TheCaptain,
}

var LoveFirstMate = LoveAndHateCard{
	IsLove:    true,
	Character: FirstMate,
}

var HateFirstMate = LoveAndHateCard{
	IsLove:    false,
	Character: FirstMate,
}

var LoveFrenchy = LoveAndHateCard{
	IsLove:    true,
	Character: Frenchy,
}

var HateFrenchy = LoveAndHateCard{
	IsLove:    false,
	Character: Frenchy,
}

var LoveDrHarter = LoveAndHateCard{
	IsLove:    true,
	Character: DrHarter,
}

var HateDrHarter = LoveAndHateCard{
	IsLove:    false,
	Character: DrHarter,
}

var LoveMadameWong = LoveAndHateCard{
	IsLove:    true,
	Character: MadameWong,
}

var HateMadameWong = LoveAndHateCard{
	IsLove:    false,
	Character: MadameWong,
}
