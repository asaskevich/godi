godi
===========
[![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/asaskevich/godi?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge) [![GoDoc](https://godoc.org/github.com/asaskevich/godi?status.png)](https://godoc.org/github.com/asaskevich/godi)
[![Build Status](https://travis-ci.org/asaskevich/godi.svg?branch=master)](https://travis-ci.org/asaskevich/godi)
[![Coverage](https://codecov.io/gh/asaskevich/godi/branch/master/graph/badge.svg)](https://codecov.io/gh/asaskevich/godi) [![Go Report Card](https://goreportcard.com/badge/github.com/asaskevich/godi)](https://goreportcard.com/report/github.com/asaskevich/godi) 

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
To start you need to create structs (like services) or factory functions.
Be sure to use tag `godi` on struct fields that will be used for the dependency injection.

```go
type Repository struct {
	Driver CustomDriver `godi:"autowire"`
    Print Printer `godi:"autowire"`
}

type CustomDriver struct {
	Name string
}

type Printer struct {
    // fields
}
...

c := Container{}
c.New()
c.RegisterService(CustomDriver{Name:"my_driver"})
c.RegisterFactory(func () Printer {
    // init Printer
    return Printer{}
})

service, err := c.ConstructService(Repository{})

println(service.(Repository).Driver.Name) // "my_driver"
```

#### List of functions and types

```go
type Container
    func (c *Container) ConstructService(item interface{}) (interface{}, error)
    func (c *Container) GetFactory(factoryType string) (func() interface{}, error)
    func (c *Container) GetService(serviceType string) (interface{}, error)
    func (c *Container) HasFactory(service interface{}) (res bool)
    func (c *Container) HasService(item interface{}) (res bool)
    func (c *Container) New()
    func (c *Container) RegisterFactory(service func() interface{})
    func (c *Container) RegisterService(item interface{})
```

#### Tag `godi`

***Note:*** without the tag inner structures will not be created

- `-` - to skip processing field
- `autowire` - allows building inner service automatically

#### Notes
Documentation is available here: [godoc.org](https://godoc.org/github.com/asaskevich/godi).
Full information about code coverage is also available here: [godi on gocover.io](http://gocover.io/github.com/asaskevich/godi).

#### Support
If you do have a contribution to the package, feel free to create a Pull Request or an Issue.

#### What to contribute
If you don't know what to do, there are some features and functions that need to be done

- [ ] Refactor code
- [ ] Edit docs and [README](https://github.com/asaskevich/godi/README.md): spellcheck, grammar and typo check
- [ ] Resolve [issues and bugs](https://github.com/asaskevich/godi/issues)
- [ ] Look at forks for new features and fixes
- [ ] Dependency injection by sub-type
- [ ] Resolve dependency conflicts (e.g. two types matched one common type)
- [ ] Qualifiers to define custom Id of service
- [ ] Autoscan
- [ ] Separate services into repositores/components/services/controllers
- [ ] Autowire string/int properties from files


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fasaskevich%2Fgodi.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fasaskevich%2Fgodi?ref=badge_large)
