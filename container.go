package godi

import (
	"fmt"
	"reflect"
)

func (c *Container) New() {
	c.services = make([]serviceWrapper, 0)
}

func (c *Container) RegisterFactory(service func() interface{}) {
	serviceType := reflect.ValueOf(service).Type().Out(0).String()

	c.services = append(c.services, serviceWrapper{
		Type:     serviceType,
		IsFunc:   true,
		Instance: service(),
	})
}

func (c *Container) RegisterService(item interface{}) {
	componentType := reflect.TypeOf(item).String()

	c.services = append(c.services, serviceWrapper{
		Type:     componentType,
		IsFunc:   false,
		Instance: item,
	})
}

func (c *Container) HasFactory(service interface{}) (res bool) {
	res = false
	t := reflect.ValueOf(service).Type().Out(0).String()

	for _, v := range c.services {
		res = res || (v.Type == t && v.IsFunc == true)
	}

	return
}

func (c *Container) HasService(item interface{}) (res bool) {
	res = false
	t := reflect.TypeOf(item).String()

	for _, v := range c.services {
		res = res || (v.Type == t && v.IsFunc == false)
	}

	return
}

func (c *Container) GetFactory(factoryType string) (func() interface{}, error) {
	res := func() interface{} {
		return nil
	}

	if factoryType == "" {
		return nil, nil
	}

	for _, v := range c.services {
		if v.Type == factoryType && v.IsFunc == true {
			if res != nil {
				return nil, fmt.Errorf("two or more factories with same implementation; got %v", factoryType)
			}
			res = v.Instance.(func() interface{})
		}
	}

	return res, nil
}

func (c *Container) GetService(serviceType string) (interface{}, error) {
	var res interface{}

	for _, v := range c.services {
		if v.Type == serviceType && v.IsFunc == false {
			if res != nil {
				return nil, fmt.Errorf("two or more services with same type; got %v", serviceType)
			}
			res = v.Instance
		}
	}

	return res, nil
}

func (c *Container) ConstructService(item interface{}) (interface{}, error) {
	if item == nil {
		return nil, fmt.Errorf("unable to construct service from null")
	}
	val := reflect.ValueOf(item)
	if val.Kind() == reflect.Interface || val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	// we only accept structs
	if val.Kind() != reflect.Struct && val.Kind() != reflect.Interface {
		return nil, fmt.Errorf("function only accepts structs; got %s", val.Kind())
	}
	rv := reflect.ValueOf(item)
	ret := reflect.New(rv.Type())

	for i := 0; i < val.NumField(); i++ {
		// copy value
		typeField := val.Type().Field(i)
		ret.Elem().FieldByName(typeField.Name).Set(val.Field(i))
		// private field
		if typeField.PkgPath != "" {
			continue
		}

		tag := typeField.Tag.Get(tagName)
		if tag != "-" && tag != "" {
			lookup := typeField.Type.String()
			service, err1 := c.GetService(lookup)
			factory, err2 := c.GetFactory(lookup)
			if err1 != nil && err2 != nil {
				return nil, fmt.Errorf("unable to find service of factory for type %v", lookup)
			}
			if service != nil {
				value, err := c.ConstructService(service)
				if err != nil {
					return nil, fmt.Errorf("unable to construct inner service of type %v", lookup)
				}
				val := reflect.New(typeField.Type).Elem()
				val.Set(reflect.ValueOf(value))
				ret.Elem().FieldByName(typeField.Name).Set(val)
			} else {
				value, err := c.ConstructService(factory())
				if err != nil {
					return nil, fmt.Errorf("unable to construct inner service of type %v", lookup)
				}
				val := reflect.New(typeField.Type).Elem()
				val.Set(reflect.ValueOf(value))
				ret.Elem().FieldByName(typeField.Name).Set(val)
			}
		}
	}
	retValue := ret.Elem().Interface()
	c.RegisterService(retValue)
	return retValue, nil
}
