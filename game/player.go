package game

import "math/rand"

type player struct {
	name   string
	score  byte
	tank   *tank
	number byte
}

func (g *game) loadPlayer(name string, first bool) *player {
	i := byte(0)
	if first == true {
		i = byte(rand.Int31n(maxPlayers))
	}
	for ; g.players[i] != nil && i < maxPlayers-1; i++ {
	}

	t := g.loadTank(i, name)
	p := &player{
		name:  name,
		score: 0,
		tank:  t,
	}
	g.players[i] = p
	return p
}

func (g *game) updatePlayer(p *player, direction Direction, moves bool) {
	g.updateTank(p.tank, direction, moves)
}

func (g *game) incrementScore(p *player) {
	p.score++
}
