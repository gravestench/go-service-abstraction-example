package abstract

type InputService interface {
	IsKeyDown(keyCode ...int) bool
	IsKeyPressed(keyCode ...int) bool
	IsKeyUp(keyCode ...int) bool
	IsKeyReleased(keyCode ...int) bool
	
	IsMouseButtonDown(buttonCode ...int) bool
}
