package abstract

// DependencyResolver describes any service which resolves its own dependencies
// using the ServiceManager. The DependencyResolver yields a slice pointers
// which are used to determine if dependency resolution is complete.
type DependencyResolver interface {
	Dependencies() []interface{}
	ResolveDependencies(manager ServiceManager)
}
