package domain

import "fmt"

type Navigation struct {
	HasBird    bool        // Отметка о прибавлении птицы (чайки)
	RemoveBird bool        // Отметка о снятии птицы (чайки)
	Overboard  []Character // Персонажи за бортом
	Thirsty    []Character // Персонажи, испытывающие жажду

	ThirstyOnFight      bool // Дополнительная жажда за драку
	ThirstyOnNavigation bool // Дополнительная жажда за навигацию
}

func (n Navigation) CardType() CardType {
	return NavigationCardType
}

func (n Navigation) Title() string {
	var overboardNames []string
	for _, character := range n.Overboard {
		overboardNames = append(overboardNames, character.Name)
	}

	var thirstyNames []string
	for _, character := range n.Thirsty {
		thirstyNames = append(thirstyNames, character.Name)
	}

	var thirstyOnNavigation = "-"
	if n.ThirstyOnNavigation {
		thirstyOnNavigation = "+"
	}

	var thirstyOnFight = "-"
	if n.ThirstyOnFight {
		thirstyOnFight = "+"
	}

	var birdText = "-"
	if n.HasBird {
		birdText = "+1"
	} else if n.RemoveBird {
		birdText = "-1"
	}

	return fmt.Sprintf("ЗА БОРТОМ (%v), ЖАЖДА (%v) | ГР: %v | ДРК: %v | ПТИЦА: %v",
		overboardNames, thirstyNames, thirstyOnNavigation, thirstyOnFight, birdText)
}

func GetNavigationAll() []Navigation {
	// For using as a shortcut only
	var noCharacter []Character

	return []Navigation{
		// Страница 9 из PDF файла с картами
		{
			HasBird:             false,
			RemoveBird:          false,
			Overboard:           GetCharacters(TheCaptain),
			Thirsty:             GetCharacters(Frenchy),
			ThirstyOnFight:      false,
			ThirstyOnNavigation: false,
		},
		{
			HasBird:             true,
			RemoveBird:          false,
			Overboard:           noCharacter,
			Thirsty:             noCharacter,
			ThirstyOnFight:      true,
			ThirstyOnNavigation: false,
		},
		{
			HasBird:             true,
			RemoveBird:          false,
			Overboard:           GetCharacters(MadameWong),
			Thirsty:             GetCharactersAllExcept(SirStephen, MadameWong, LadyLauren, TheKid),
			ThirstyOnFight:      false,
			ThirstyOnNavigation: true,
		},
		{
			HasBird:             false,
			RemoveBird:          false,
			Overboard:           GetCharacters(FirstMate),
			Thirsty:             GetCharacters(TheCaptain, Frenchy),
			ThirstyOnFight:      true,
			ThirstyOnNavigation: true,
		},
		{
			HasBird:             true,
			RemoveBird:          false,
			Overboard:           GetCharacters(Frenchy),
			Thirsty:             GetCharacters(TheCaptain, FirstMate, SirStephen),
			ThirstyOnFight:      false,
			ThirstyOnNavigation: false,
		},
		{
			HasBird:             false,
			RemoveBird:          false,
			Overboard:           GetCharacters(SirStephen),
			Thirsty:             GetCharacters(TheCaptain, FirstMate, Frenchy),
			ThirstyOnFight:      true,
			ThirstyOnNavigation: true,
		},
		{
			HasBird:             false,
			RemoveBird:          true,
			Overboard:           GetCharacters(LadyLauren),
			Thirsty:             GetCharactersAll(),
			ThirstyOnFight:      false,
			ThirstyOnNavigation: true,
		},
		// Страница 11 из PDF файла с картами
		{
			HasBird:             false,
			RemoveBird:          false,
			Overboard:           GetCharacters(SirStephen),
			Thirsty:             GetCharacters(TheCaptain, LadyLauren),
			ThirstyOnFight:      true,
			ThirstyOnNavigation: false,
		},
		{
			HasBird:             false,
			RemoveBird:          false,
			Overboard:           GetCharacters(FirstMate),
			Thirsty:             GetCharacters(TheCaptain, FirstMate),
			ThirstyOnFight:      false,
			ThirstyOnNavigation: true,
		},
		{
			HasBird:             false,
			RemoveBird:          false,
			Overboard:           GetCharacters(TheKid),
			Thirsty:             GetCharacters(TheCaptain, FirstMate, Frenchy, LadyLauren),
			ThirstyOnFight:      false,
			ThirstyOnNavigation: false,
		},
		{
			HasBird:             true,
			RemoveBird:          false,
			Overboard:           GetCharacters(SirStephen),
			Thirsty:             GetCharacters(TheCaptain, TheKid),
			ThirstyOnFight:      false,
			ThirstyOnNavigation: true,
		},
		{
			HasBird:             false,
			RemoveBird:          false,
			Overboard:           GetCharacters(Frenchy),
			Thirsty:             GetCharacters(TheCaptain, FirstMate, TheKid),
			ThirstyOnFight:      false,
			ThirstyOnNavigation: true,
		},
		{
			HasBird:             false,
			RemoveBird:          false,
			Overboard:           GetCharacters(Frenchy),
			Thirsty:             GetCharacters(TheCaptain, FirstMate, LadyLauren),
			ThirstyOnFight:      true,
			ThirstyOnNavigation: false,
		},
		{
			HasBird:             false,
			RemoveBird:          false,
			Overboard:           GetCharacters(SirStephen),
			Thirsty:             GetCharacters(TheCaptain, SirStephen),
			ThirstyOnFight:      false,
			ThirstyOnNavigation: false,
		},
		// Пздц мощная карта
		{
			HasBird:             true,
			RemoveBird:          false,
			Overboard:           GetCharactersAll(),
			Thirsty:             GetCharactersAll(),
			ThirstyOnFight:      true,
			ThirstyOnNavigation: true,
		},
		// Страница 13 из PDF файла с картами
		{
			HasBird:             false,
			RemoveBird:          false,
			Overboard:           GetCharacters(TheKid),
			Thirsty:             GetCharactersAllExcept(TheKid),
			ThirstyOnFight:      false,
			ThirstyOnNavigation: true,
		},
		{
			HasBird:             false,
			RemoveBird:          false,
			Overboard:           GetCharacters(TheKid),
			Thirsty:             GetCharactersAllExcept(LadyLauren),
			ThirstyOnFight:      true,
			ThirstyOnNavigation: false,
		},
		{
			HasBird:             true,
			RemoveBird:          false,
			Overboard:           GetCharacters(TheKid),
			Thirsty:             GetCharactersAllExcept(LadyLauren),
			ThirstyOnFight:      true,
			ThirstyOnNavigation: true,
		},
		{
			HasBird:             false,
			RemoveBird:          false,
			Overboard:           GetCharacters(TheCaptain),
			Thirsty:             GetCharacters(SirStephen),
			ThirstyOnFight:      true,
			ThirstyOnNavigation: true,
		},
		{
			HasBird:             false,
			RemoveBird:          false,
			Overboard:           GetCharacters(TheCaptain),
			Thirsty:             GetCharacters(TheCaptain),
			ThirstyOnFight:      false,
			ThirstyOnNavigation: false,
		},
		{
			HasBird:             false,
			RemoveBird:          false,
			Overboard:           GetCharacters(TheCaptain),
			Thirsty:             GetCharacters(TheCaptain, FirstMate, Frenchy, SirStephen),
			ThirstyOnFight:      true,
			ThirstyOnNavigation: true,
		},
		{
			HasBird:             true,
			RemoveBird:          false,
			Overboard:           GetCharacters(TheCaptain),
			Thirsty:             GetCharacters(FirstMate),
			ThirstyOnFight:      true,
			ThirstyOnNavigation: false,
		},
		{
			HasBird:             true,
			RemoveBird:          false,
			Overboard:           GetCharacters(FirstMate),
			Thirsty:             GetCharacters(TheKid),
			ThirstyOnFight:      true,
			ThirstyOnNavigation: false,
		},
		// Страница 15 из PDF файла с картами
		{
			HasBird:             false,
			RemoveBird:          false,
			Overboard:           GetCharacters(FirstMate, LadyLauren),
			Thirsty:             GetCharacters(LadyLauren),
			ThirstyOnFight:      false,
			ThirstyOnNavigation: false,
		},
		{
			HasBird:             false,
			RemoveBird:          false,
			Overboard:           GetCharacters(DrHarter),
			Thirsty:             GetCharacters(FirstMate, DrHarter),
			ThirstyOnFight:      true,
			ThirstyOnNavigation: false,
		},
		{
			HasBird:             true,
			RemoveBird:          false,
			Overboard:           GetCharacters(DrHarter),
			Thirsty:             GetCharacters(TheCaptain, FirstMate),
			ThirstyOnFight:      true,
			ThirstyOnNavigation: true,
		},
		{
			HasBird:             false,
			RemoveBird:          false,
			Overboard:           GetCharacters(MadameWong),
			Thirsty:             GetCharactersAllExcept(TheCaptain, DrHarter, MadameWong, TheKid),
			ThirstyOnFight:      false,
			ThirstyOnNavigation: false,
		},
		{
			HasBird:             false,
			RemoveBird:          false,
			Overboard:           GetCharacters(DrHarter),
			Thirsty:             GetCharacters(SirStephen, Frenchy),
			ThirstyOnFight:      false,
			ThirstyOnNavigation: false,
		},
		{
			HasBird:             false,
			RemoveBird:          false,
			Overboard:           GetCharacters(DrHarter),
			Thirsty:             GetCharacters(TheCaptain, DrHarter),
			ThirstyOnFight:      false,
			ThirstyOnNavigation: true,
		},
	}
}
