package light

type Light interface {
	Shine()
	Charge() Light
}
