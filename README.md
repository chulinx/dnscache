# dnscache for  go

> Resolution dns resolver time to long in Go


# Usage

```go
var resolve = New()

func ResolverBaiduUseCache(host string) ([]string, error) {
    ips, err := resolve.LookupHosts(host)
    return ips, err
}
```

# Install 

```shell
go get github.com/chulinx/dnscache
```