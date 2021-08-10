
# how to publish you go module and make it found by go-proxy
1. set release tag, and release the module in git repo

2. run as below from another go-module (tag as v0.0.1)
```
GOPROXY=goproxy.cn go list -m github.com/13ye/test-go-mod@v0.0.1
go get github.com/13ye/test-go-mod@v0.0.1
```
