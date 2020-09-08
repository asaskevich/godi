godi
===========
[![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/asaskevich/godi?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge) [![GoDoc](https://godoc.org/github.com/asaskevich/godi?status.png)](https://godoc.org/github.com/asaskevich/godi)
[![Build Status](https://travis-ci.org/asaskevich/godi.svg?branch=master)](https://travis-ci.org/asaskevich/godi)
[![Coverage](https://codecov.io/gh/asaskevich/godi/branch/master/graph/badge.svg)](https://codecov.io/gh/asaskevich/godi) [![Go Report Card](https://goreportcard.com/badge/github.com/asaskevich/godi)](https://goreportcard.com/report/github.com/asaskevich/godi) [![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fasaskevich%2Fgodi.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fasaskevich%2Fgodi?ref=badge_shield)


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

## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fasaskevich%2Fgodi.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fasaskevich%2Fgodi?ref=badge_large)