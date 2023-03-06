package abstract

// Service is an abstraction for what a service is
type Service interface {
	Init(possibleDependencies *[]interface{})
	Name() string
}
