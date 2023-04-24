package abstract

// Service is an abstraction for what a service is
type Service interface {
	Init(ServiceManager)
	Name() string
}

// ServiceManager describes a service that manages other services
type ServiceManager interface {
	AddServices(...Service)
	AddService(Service)
	RemoveService(Service)
	Services() *[]interface{}
	Quit()
}
