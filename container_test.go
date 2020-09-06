package godi

import "testing"

func TestContainer_New(t *testing.T) {
	c := Container{}
	c.New()

	if &c == nil {
		t.Fatalf("container is not created; c = %v", c)
	}

	if &GlobalContainer == nil {
		t.Fatalf("global container is not created; c = %v", c)
	}

	if len(c.services) != 0 {
		t.Fatalf("the list of services is not empty; c.services = %v", c.services)
	}
}

func TestContainer_ConstructService(t *testing.T) {
	c := Container{}
	c.New()
	c.RegisterService(CustomDriver{Name: "test"})

	service, err := c.ConstructService(Repository{})

	if err != nil {
		t.Fatalf("error is not empty; err = %v", err)
	}

	if &service == nil {
		t.Fatal("service is not created")
	}

	if service.(Repository).Driver.Name != "test" {
		t.Fatal("inner service is not created")
	}
}
