package example

var resolve = dnscache.New()

func ResolverBaiduUseCache(host string) ([]string, error) {
	ips, err := resolve.LookupHosts(host)
	return ips, err
}
