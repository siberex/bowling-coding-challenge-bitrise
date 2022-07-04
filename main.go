package main

import (
	"fmt"

	"github.com/siberex/go_bowling/bowling"
)

func main() {
	{
		// simple game
		g := bowling.Game{}

		// 5
		g.Roll(2)
		g.Roll(3)

		// 5
		g.Roll(2)
		g.Roll(3)

		// 3
		g.Roll(1)
		g.Roll(2)

		fmt.Println(g.Score()) // 13
	}

	{
		// spare
		g := bowling.Game{}

		// 10 + 5
		g.Roll(8)
		g.Roll(2)

		// 6
		g.Roll(5)
		g.Roll(1)

		// 3
		g.Roll(1)
		g.Roll(2)

		fmt.Println(g.Score()) // 24
	}

	{
		// strike
		g := bowling.Game{}

		// 10 + 6
		g.Roll(10)

		// 6
		g.Roll(5)
		g.Roll(1)

		// 1
		g.Roll(1)

		fmt.Println(g.Score()) // 23
	}
}
