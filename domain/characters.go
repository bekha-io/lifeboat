package domain

import "math/rand"

type Character struct {
	Name          string
	Size          int // Damage and health both
	SurvivalValue int
	Ability       CharacterAbility

	Damage int
}

func (c Character) String() string {
	return c.Title()
}

func (c Character) CardType() CardType {
	return CharacterCardType
}

func (c Character) Title() string {
	return c.Name
}

func (c *Character) IsAlive() bool {
	return c.Damage > c.Size
}

func (c *Character) HasConsciousness() bool {
	return c.Damage != c.Size
}

// TheKid Шкет
var TheKid = Character{
	Name:          "Шкет",
	Size:          3,
	SurvivalValue: 9,
	Ability:       CanStealFaceDownCardsAbility,
}

// LadyLauren Миледи
var LadyLauren = Character{
	Name:          "Миледи",
	Size:          4,
	SurvivalValue: 8,
	Ability:       DoublePointsForJewelAbility,
}

// SirStephen Сноб
var SirStephen = Character{
	Name:          "Сноб",
	Size:          5,
	SurvivalValue: 7,
	Ability:       DoublePointsForPaintingsAbility,
}

// TheCaptain Капитан
var TheCaptain = Character{
	Name:          "Капитан",
	Size:          7,
	SurvivalValue: 5,
	Ability:       DoublePointsForCashAbility,
}

// FirstMate Боцман
var FirstMate = Character{
	Name:          "Боцман",
	Size:          8,
	SurvivalValue: 4,
	Ability:       HeIsJustPlainBigAbility,
}

// Frenchy Черпак
var Frenchy = Character{
	Name:          "Черпак",
	Size:          6,
	SurvivalValue: 6,
	Ability:       ExcellentSwimmerAbility,
}

// DrHarter Доктор Хартер
var DrHarter = Character{
	Name:          "Доктор Хартер",
	Size:          4,
	SurvivalValue: 8,
	Ability:       ReusableMedicalAbility,
}

// MadameWong Мадам Вонг
var MadameWong = Character{
	Name:          "Мадам Вонг",
	Size:          3,
	SurvivalValue: 9,
	Ability:       GainWaterBenefitAsWellAbility,
}

var allCharacters = []Character{
	TheKid, LadyLauren, SirStephen, TheCaptain, FirstMate, Frenchy, DrHarter, MadameWong,
}

func GetCharactersAllExcept(characters ...Character) []Character {
	var chs []Character
	for _, ch := range allCharacters {
		for _, chh := range characters {
			if ch.Name == chh.Name {
				continue
			}
		}
		chs = append(chs, ch)
	}
	return chs
}

func GetCharactersAll() []Character {
	return allCharacters
}

func GetCharacters(characters ...Character) []Character {
	var chs []Character
	for _, ch := range characters {
		chs = append(chs, ch)
	}
	return chs
}

func getCharactersRandom(n int) []*Character {
	chs := GetCharactersAll()
	rand.Shuffle(len(chs), func(i, j int) {
		chs[i], chs[j] = chs[j], chs[i]
	})

	var playingChs []*Character
	for i := 1; i >= n; i++ {
		playingChs = append(playingChs, &chs[i])
	}

	return playingChs
}
