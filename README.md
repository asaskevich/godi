godi
===========
Dependency injection in Golang at runtime.

#### Installation
Make sure that Go is installed on your computer.
Type the following command in your terminal:

	go get github.com/asaskevich/godi

or you can get specified release of the package with `gopkg.in`:

	go get gopkg.in/asaskevich/godi.v1.0.0

After it the package is ready to use.


#### Import package in your project
Add following line in your `*.go` file:
```go
import "github.com/asaskevich/godi"
```

#### Example of usage
```go
type Repository struct {
	Driver CustomDriver `godi:"autowire"`
}

type CustomDriver struct {
	Name string
}
...

c := Container{}
c.New()
c.RegisterService(CustomDriver{Name:"my_driver"})

service, err := c.ConstructService(Repository{})

println(service.(Repository))
```