package domain

type Fight struct {
	AttackingPlayers []Player
	DefensivePlayers []Player
}

// AttackingSize общий урон нападающих
func (f Fight) AttackingSize() int {
	var s int

	for _, pl := range f.AttackingPlayers {
		s += pl.Character.Size
	}

	return s
}

// DefensiveSize общий урон защищающихся
func (f Fight) DefensiveSize() int {
	var s int

	for _, pl := range f.DefensivePlayers {
		s += pl.Character.Size
	}

	return s
}
