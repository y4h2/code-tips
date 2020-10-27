package naive

var instance *LoadBalancer = nil

type LoadBalancer struct {
	serverList []string
}

func (lb *LoadBalancer) AddServer(server string) {
	lb.serverList = append(lb.serverList, server)
}

func (lb *LoadBalancer) RemoveServer(server string) {
	var pos = -1
	for i, s := range lb.serverList {
		if s == server {
			pos = i
			break
		}
	}
	if pos >= 0 {
		lb.serverList = append(lb.serverList[:pos], lb.serverList[pos+1:]...)
	}
}

func newLoadBalancer() *LoadBalancer {
	return &LoadBalancer{
		serverList: []string{},
	}
}

func GetLoadBalancer() *LoadBalancer {
	if instance == nil {
		instance = newLoadBalancer()
	}

	return instance
}
