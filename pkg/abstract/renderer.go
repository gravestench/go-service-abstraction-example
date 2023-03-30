package abstract

type Renderer interface {
	Name() string
	IsActive() bool
	Render()
}

type RenderService interface {
	Initialized() bool
	WindowWidth() int
	WindowHeight() int
	Close()
}
