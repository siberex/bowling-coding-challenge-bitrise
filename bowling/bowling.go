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

		if i >= 2 && (i%2 == 0) && (g.rolls[i-2]+g.rolls[i-1] == 10) {
			sum += n
		}
	}

	return sum
}
