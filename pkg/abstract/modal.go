package abstract

type Mode int

type Modal interface {
	Mode() Mode
	SetMode(Mode)
}

type ModeService interface {
	Modals() []Modal
	SetMode(Mode)
}
