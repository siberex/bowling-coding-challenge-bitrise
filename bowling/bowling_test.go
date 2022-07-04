package bowling_test

import (
	"testing"

	"github.com/siberex/go_bowling/bowling"
)

func TestScore(t *testing.T) {
	testCases := []struct {
		desc    string
		rollAll func(g *bowling.Game)
		want    int
	}{
		{
			desc: "simple game",
			rollAll: func(g *bowling.Game) {
				g.Roll(1)
				g.Roll(2)

				g.Roll(3)
				g.Roll(1)

				g.Roll(0)
				g.Roll(1)
			},
			want: 3 + 4 + 1,
		},
		{
			desc: "spare game",
			rollAll: func(g *bowling.Game) {
				g.Roll(1)
				g.Roll(2)

				g.Roll(3)
				g.Roll(7)

				g.Roll(2)
				g.Roll(1)

			},
			want: 3 + 10 + 2 + 3,
		},
		{
			desc: "strike game",
			rollAll: func(g *bowling.Game) {
				g.Roll(1)
				g.Roll(2)

				g.Roll(10)

				g.Roll(2)
				g.Roll(1)
			},
			want: 3 + 10 + 2 + 1 + 3, // 19
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			g := bowling.Game{}
			tc.rollAll(&g)
			got := g.Score()
			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
