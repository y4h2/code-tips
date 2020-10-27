package check

import (
	"sync"
	"sync/atomic"
)

var instance *LoadBalancer = nil
var mu sync.Mutex

var inited uint32 = 0

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

/* will have race condition
在instance == nil里面加锁, 只会让goroutine在if的block里面等待
最后依然会运行 instance = newLoadBalancer()
*/
func GetLoadBalancer() *LoadBalancer {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = newLoadBalancer()
		}
	}

	return instance
}

// 模拟sync.Once的解决方案
func GetLoadBalancerFixed() *LoadBalancer {
	if atomic.LoadUint32(&inited) == 0 {
		mu.Lock()
		defer mu.Unlock()
		if atomic.LoadUint32(&inited) == 0 {
			instance = newLoadBalancer()
			atomic.StoreUint32(&inited, uint32(1))
		}
	}

	return instance
}
