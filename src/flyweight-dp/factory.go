package flyweight_dp

import "fmt"

const (
	//AdminType terrorist right type
	AdminType = "admin"
	//UserType terrorist right type
	UserType = "user"
)

var (
	rightFactoryInstance = &DressFactory{
		playerRightMap: make(map[string]IRight),
	}
)

type DressFactory struct {
	playerRightMap map[string]IRight
}

func getRightFactoryInstance() *DressFactory {
	return rightFactoryInstance
}

func (d *DressFactory) getRightByType(rightType string) (IRight, error) {
	if d.playerRightMap[rightType] != nil {
		return d.playerRightMap[rightType], nil
	}
	switch rightType {
	case AdminType:
		d.playerRightMap[rightType] = newAdminRight()
		return d.playerRightMap[rightType], nil
	case UserType:
		d.playerRightMap[rightType] = newUserRight()
		return d.playerRightMap[rightType], nil
	}
	return nil, fmt.Errorf("Wrong rightType  passed")
}
