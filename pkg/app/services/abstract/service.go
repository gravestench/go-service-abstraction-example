package abstract

type Service interface {
	Init(possibleDependencies *[]interface{})
	Name() string
}
