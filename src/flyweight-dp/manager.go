package flyweight_dp

import "go-dps/pkg/common/array"

type (
	manager struct {
		playerRights map[string]*Player
		players      map[string][]string
	}
)

func newManager() *manager {
	return &manager{
		playerRights: make(map[string]*Player),  // 类型->权限
		players:      make(map[string][]string), // 类型 ->用户
	}
}

func (m *manager) addPlayers(rightType string, users ...string) {
	for _, user := range users {
		if _, ok := m.players[rightType]; ok {
			if array.StrFind(m.players[rightType], user) {
				continue
			}
			m.players[rightType] = append(m.players[rightType], user)
			continue
		}
		m.players[rightType] = append(m.players[rightType], user)
		player := newPlayer(rightType)
		m.playerRights[rightType] = player
	}
	return
}

func (m *manager) getUserAndRight(rightTypes ...string) map[string]*Player {
	var playerAndRight = make(map[string]*Player)
	for _, rightType := range rightTypes {
		if _, ok := m.playerRights[rightType]; !ok {
			continue
		}
		for _, user := range m.players[rightType] {
			if _, flag := playerAndRight[user]; flag {
				continue
			}
			playerAndRight[user] = m.playerRights[rightType]
		}
	}
	return playerAndRight
}
