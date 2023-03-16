package domain

import (
	"fmt"
	"testing"
)

func TestNewBoatSeats_CorrectOrder(t *testing.T) {
	chs := []*Character{
		&TheKid, &TheCaptain, &LadyLauren, &Frenchy,
	}
	bss := NewBoatSeats(chs...)

	for _, bs := range bss {
		fmt.Println(bs.Character.Size)
	}

	TheKid.Size = -1

	for _, bs := range bss {
		fmt.Println(bs.Character.Size)
	}
}
