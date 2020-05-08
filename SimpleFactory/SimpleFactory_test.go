package SimpleFactory

import "testing"

func TestSimpleFactory(t *testing.T){


	b := &Factory{need:'b'}
	b.factory().produce()

	c := &Factory{need:'c'}
	c.factory().produce()
}