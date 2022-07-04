package bowling

const PINS_PER_FRAME = 10
const ATTEMPTS_PER_FRAME = 2

type Game struct {
	rolls  []int // deprecated
	frames []Frame
}

type Frame struct {
	attempts []int
}

func (fr *Frame) IsStrike() bool {
	return len(fr.attempts) == 1 && fr.GetScore() == PINS_PER_FRAME
}

func (fr *Frame) IsSpare() bool {
	return len(fr.attempts) == ATTEMPTS_PER_FRAME && fr.GetScore() == PINS_PER_FRAME
}

// false = could not make any more attemts in this Frame
// and should proceed to the next frame
func (fr *Frame) Attempt(n int) bool {
	if n > PINS_PER_FRAME {
		return false
	}

	// Could not score more pins than left after the first attempt
	if len(fr.attempts) == 1 && fr.attempts[0]+n > PINS_PER_FRAME {
		return false
	}

	if fr.IsStrike() {
		return false
	}

	// Final attempt already scored in this Frame
	if len(fr.attempts) >= ATTEMPTS_PER_FRAME {
		return false
	}

	fr.attempts = append(fr.attempts, n)
	return true
}

func (fr *Frame) GetScore() int {
	total := 0
	for _, v := range fr.attempts {
		total += v
	}
	return total
}

func (g *Game) Roll(n int) {
	g.rolls = append(g.rolls, n)

	// Add new frame and immediately store attempt there
	var fr Frame
	_ = fr.Attempt(n)

	if len(g.frames) == 0 {
		// First attempt, first frame
		g.frames = append(g.frames, fr)
	} else {
		// Get last frame
		lastFrame := &g.frames[len(g.frames)-1]

		// If run out of attempts, add new frame
		if !lastFrame.Attempt(n) {
			g.frames = append(g.frames, fr)
		}
	}
}

func (g *Game) Score() int {
	//return g.scoreFromRolls() // simple solution

	score := 0

	for i, fr := range g.frames {
		score += fr.GetScore()

		if i > 0 {
			prevFrame := g.frames[i-1]

			if prevFrame.IsSpare() {
				// Double Score for Attempt that comes after Spare
				score += fr.attempts[0]
			}
			if prevFrame.IsStrike() {
				// Double score for Frame that comes after Strike
				score += fr.GetScore()
			}
		}
	}

	return score
}

func (g *Game) scoreFromRolls() int {
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
