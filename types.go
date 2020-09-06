package godi

type serviceWrapper struct {
	Type     string
	IsFunc   bool
	Instance interface{}
}

// Container holds all initialized services and factories in it
type Container struct {
	// the list of registered and initialized services
	services []serviceWrapper
}

// GlobalContainer is the container that can be used as a main scope of dependency injection
// and functions/factories inside it
var GlobalContainer Container

func init() {
	GlobalContainer := Container{}
	GlobalContainer.New()
}
