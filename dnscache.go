package dnscache

import (
	"context"
	"math/rand"
	"net"
	"reflect"
	"sync"
	"time"
)

var store = make(map[string][]string)

type resolverCache struct {
	Context  context.Context
	Resolver *net.Resolver
	Time     time.Duration
	Lock     sync.RWMutex
}

func New() *resolverCache {
	ctx := context.Background()
	r := &resolverCache{
		Context: ctx,
	}
	if r.Time == 0 {
		r.Time = time.Second * 30
	}
	go func() {
		for {
			for k, v := range store {
				ips, err := r.lookup(k)
				if err != nil {
					continue
				}
				if reflect.DeepEqual(v, ips) {
					continue
				}
				r.store(k, ips)
			}
			time.Sleep(r.Time)
		}
	}()
	return r
}

func (r *resolverCache) LookupHosts(domain string) ([]string, error) {
	return r.lookup(domain)
}

func (r *resolverCache) LookupOneHost(domain string) string {
	ips, err := r.lookup(domain)
	if err != nil {
		return ""
	}
	rands := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := rands.Intn(len(ips))
	return ips[index]
}

func (r *resolverCache) lookup(domain string) ([]string, error) {
	var (
		value interface{}
		ok    bool
	)
	r.Lock.RLock()
	value, ok = store[domain]
	r.Lock.RUnlock()
	if !ok {
		return r.lookupHost(domain)
	}
	return value.([]string), nil
}

func (r *resolverCache) lookupHost(domain string) ([]string, error) {
	ips, err := r.Resolver.LookupHost(r.Context, domain)
	if err == nil {
		r.store(domain, ips)
	}
	return ips, err
}

func (r *resolverCache) store(domain string, ips []string) {
	r.Lock.Lock()
	store[domain] = ips
	r.Lock.Unlock()
}
