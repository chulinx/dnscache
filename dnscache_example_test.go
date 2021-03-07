package dnscache

import (
	"fmt"
	"sync"
	"testing"
)

const host = "www.baidu.com"

var wg sync.WaitGroup
var lock sync.RWMutex
var hosts = [3]string{"www.baidu.com", "www.github.com", "gitee.com"}
var resolveMap = map[string]string{
	"180.101.49.11":  "www.baidu.com",
	"180.101.49.12":  "www.baidu.com",
	"112.80.248.75":  "www.baidu.com",
	"112.80.248.76": "www.baidu.com",
	"13.229.188.59":  "www.github.com",
	"212.64.62.183": "gitee.com",
}

func TestResolverBaiduUseCache(t *testing.T) {
	length := 10000
	g := length * len(hosts)
	wg.Add(g)
	for i := 0; i < length; i++ {
		for _, h := range hosts {
			go func(h string) {
				defer wg.Done()
				ips, _ := ResolverBaiduUseCache(h)
				lock.RLock()
				result := resolveMap[ips[0]]
				lock.RUnlock()
				if result != h {
					fmt.Printf("H: %s=>Host: %s\n",h, result)
					t.Error("Failed")
				}
			}(h)
		}
	}
	wg.Wait()
	t.Log("Success")
}

func BenchmarkNameResolverBaiduUseCache(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := ResolverBaiduUseCache(host)
		if err == nil {
			continue
		}
	}
}
