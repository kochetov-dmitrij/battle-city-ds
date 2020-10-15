package game

type player struct {
	name   string
	score  byte
	tank   *tank
	number byte
}

func (g *game) loadPlayer(name string) *player {
	i := byte(0)
	for ; g.players[i] != nil && i < 3; i++ {
	}

	t := g.loadTank(i)
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
