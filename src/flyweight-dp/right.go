package flyweight_dp

type IRight interface {
	getRight() string
}

type AdminRight struct {
	right string
}

func (t *AdminRight) getRight() string {
	return t.right
}

func newAdminRight() *AdminRight {
	return &AdminRight{right: "admin|user"}
}

type UserRight struct {
	right string
}

func (c *UserRight) getRight() string {
	return c.right
}

func newUserRight() *UserRight {
	return &UserRight{right: "user"}
}
