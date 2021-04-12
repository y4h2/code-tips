package traffic

type TrafficLight interface {
	StartTimer()
}

type GreenLight struct {
}

func (light *GreenLight) StartTimer() {

}

type YellowLight struct {
}
type RedLight struct {
}

type TrafficSystem struct {
	light TrafficLight
}

func (t *TrafficSystem) SetLight(light TrafficLight) {
	t.light = light
}

func (t *TrafficSystem) Run() {
	t.light.StartTimer()
	// move to next
}
