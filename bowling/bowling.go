package bowling

type Game struct {
	rolls []int
}

func (g *Game) Roll(n int) {
	g.rolls = append(g.rolls, n)
}

func (g *Game) Score() int {
	sum := 0
	for i, n := range g.rolls {
		sum += n

		if i >= 1 && g.rolls[i-1] == 10 {
			// Previous Frame was a Strike
			sum += n
		} else if i >= 2 && g.rolls[i-2] == 10 {
			// Previous Frame was a Strike
			sum += n
		} else if i >= 2 && (i%2 == 0) && (g.rolls[i-2]+g.rolls[i-1] == 10) {
			// Previous Frame was a Spare
			sum += n
		}
	}

	return sum
}
