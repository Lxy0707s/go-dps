package flyweight_dp

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	// 例如，存储，处理map[string]map[string]interface 的数据，代表用户，用户类型，用户权限
	// 不使用享元模式，则如果用户拥有100W用户量级，则后面的每一个用户类型， 用户权限会浪费大量空间
	// 使用享元模式，则将分为，用户类型-用户，用户类型-权限，分别使用缓存存储，优先使用缓存，没有缓存才会创建对应的权限对象，节省空间
	//   1. 权限工厂，负责缓存处理 用户类型-权限
	//   2. 管理中心，负责缓存处理 用户类型-用户，同时检测是否需要权限工厂创建权限
	mg := newManager()

	//Add Admin
	mg.addPlayers(AdminType, "admin1", "admin2")

	//Add User
	mg.addPlayers(UserType, "user1", "user2", "user3")

	//Get Player type and Right
	for user, player := range mg.getUserAndRight(AdminType, UserType) {
		fmt.Printf("User: %s, RightType: %s,  Right: %s\n", user, player.playerType, player.right.getRight())
	}
}
