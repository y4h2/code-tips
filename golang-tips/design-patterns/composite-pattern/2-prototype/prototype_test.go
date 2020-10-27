package prototype

import "testing"

func TestLeafImplement(t *testing.T) {
	var _ Component = &Leaf{}
}

func TestCompositeImplement(t *testing.T) {
	var _ Component = &Composite{}
}
