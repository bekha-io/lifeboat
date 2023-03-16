package domain

type CharacterAbility struct {
	Description string
	Callable    func()
}

// CanStealFaceDownCardsAbility способность Шкета
var CanStealFaceDownCardsAbility = CharacterAbility{
	"Может забрать любую закрытую карту без последствий",
	func() {

	},
}

// DoublePointsForJewelAbility способность Миледи
var DoublePointsForJewelAbility = CharacterAbility{
	Description: "Двойной бонус за украшения",
	Callable: func() {

	},
}

// DoublePointsForPaintingsAbility способность Сноба
var DoublePointsForPaintingsAbility = CharacterAbility{
	Description: "Двойной бонус за картины",
	Callable: func() {

	},
}

// DoublePointsForCashAbility способность Капитана
var DoublePointsForCashAbility = CharacterAbility{
	Description: "Двойной бонус за деньги",
	Callable: func() {

	},
}

// HeIsJustPlainBigAbility способность Боцмана (ее нет, на самом деле)
var HeIsJustPlainBigAbility = CharacterAbility{
	Description: "Здоровый как черт",
	Callable: func() {

	},
}

// ExcellentSwimmerAbility способность Черпака
var ExcellentSwimmerAbility = CharacterAbility{
	Description: "Хорошо плавает. Не получает ранений при падении за борт",
	Callable: func() {

	},
}

// ReusableMedicalAbility способность Доктора Хартера
var ReusableMedicalAbility = CharacterAbility{
	Description: "Не сбрасывает карту аптечки при использовании",
	Callable: func() {

	},
}

// GainWaterBenefitAsWellAbility способность Мадам Вонг
var GainWaterBenefitAsWellAbility = CharacterAbility{
	Description: "Когда любой из игроков использует воду или огненную воду, Мадам Вонг может также получить эффект от этого",
	Callable: func() {

	},
}
