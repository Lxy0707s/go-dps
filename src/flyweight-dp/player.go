package flyweight_dp

type Player struct {
	right      IRight
	playerType string
}

func newPlayer(playerType string) *Player {
	right, _ := getRightFactoryInstance().getRightByType(playerType)
	return &Player{
		playerType: playerType,
		right:      right,
	}
}
