package dnscache

var resolve = New()

func ResolverBaiduUseCache(host string) ([]string, error) {
	ips, err := resolve.LookupHosts(host)
	return ips, err
}
