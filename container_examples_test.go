package godi

func ExampleContainer_New() {
	container := Container{}
	container.New()
	// or you may use global container
	//container := GlobalContainer
}

func ExampleContainer_ConstructService() {
	c := Container{}
	c.New()
	c.RegisterService(CustomDriver{Name: "driver"})

	service, _ := c.ConstructService(Repository{})

	println(service.(Repository).Driver.Name) // "driver"
}
