package prototype

type Component interface {
	Add(c Component)
	Remove(c Component)
	GetChild(i int) Component
	Operation()
}

type Leaf struct {
}

func (l *Leaf) Add(c Component) {

}

func (l *Leaf) Remove(c Component) {

}

func (l *Leaf) GetChild(i int) Component {
	return nil
}

func (l *Leaf) Operation() {

}

type Composite struct {
	list []Component
}

func (com *Composite) Add(c Component) {

}

func (com *Composite) Remove(c Component) {

}

func (com *Composite) GetChild(i int) Component {
	return nil
}

func (com *Composite) Operation() {

}
