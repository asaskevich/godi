package godi

type Driver interface {
	Get() string
	Set(value string)
}

type Repository struct {
	Driver CustomDriver `godi:"autowire"`
}

type CustomDriver struct {
	Name string
}

func (c *CustomDriver) Get() string {
	return "CustomDriver"
}
func (c *CustomDriver) Set(value string) {

}

func GetDriver() Driver {
	return &CustomDriver{Name: "Driver"}
}

func GetCustomDriver() CustomDriver {
	return CustomDriver{Name: "CustomDriver"}
}
