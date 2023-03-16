package domain

import "sort"

// Порядок сидений от носа шлюпки к корме
var characterBoatSeatOrderMap = map[string]int{
	LadyLauren.Name: 1,
	SirStephen.Name: 2,
	TheCaptain.Name: 3,
	FirstMate.Name:  4,
	MadameWong.Name: 5,
	Frenchy.Name:    6,
	DrHarter.Name:   7,
	TheKid.Name:     8,
}

// Boat объект лодки
type Boat struct {
	Seats []*BoatSeat
}

// BoatSeat объект сидения в лодке
type BoatSeat struct {
	Character *Character
}

func (b *BoatSeat) IsTaken() bool {
	return b.Character != nil
}

func NewBoat(characters ...*Character) *Boat {
	return &Boat{
		Seats: NewBoatSeats(characters...),
	}
}

func NewBoatSeats(characters ...*Character) []*BoatSeat {
	var bss []*BoatSeat
	for _, character := range characters {
		bss = append(bss, &BoatSeat{Character: character})
	}

	sort.SliceStable(bss, func(i, j int) bool {
		return characterBoatSeatOrderMap[bss[i].Character.Name] < characterBoatSeatOrderMap[bss[j].Character.Name]
	})
	return bss
}
