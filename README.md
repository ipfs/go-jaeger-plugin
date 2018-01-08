go-jaeger-plugin
==================
Allows usage of Jaeger Bindings for Go OpenTracing API
For more information on Jaeger - https://github.com/jaegertracing/jaeger-client-go
## Setup
* NOTE: As of [2017.08.09](https://golang.org/pkg/plugin/) the plugins lib
in Go only works in Linux.

go get the code without building
```
$ go get -d github.com/ipfs/go-jaeger-plugin
```
install deps
```
$ cd $GOPATH/src/github.com/ipfs/go-jaeger-plugin/
```
```
$ gx install
```
build the plugin
```
$ cd plugin
```
```
$ make
```
## Installing
Move `jaeger-plugin.so` to `$IPFS_PATH/plugins/jaeger-plugin.so` and set it to be executable:

Make plugin executable
```
$ chmod +x jaeger-plugin.so
```
Copy the plugin to ./ipfs
```
$ cp jaeger-plugin.so $IPFS_PATH/plugins/jaeger-plugin.so
```
## Viewing Traces
```
docker run -d -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p5775:5775/udp -p6831:6831/udp -p6832:6832/udp \
  -p5778:5778 -p16686:16686 -p14268:14268 -p9411:9411 jaegertracing/all-in-one:latest
```
Open `localhost:16686` in browser
For more information on getting started with Jaeger UI
- https://jaeger.readthedocs.io/en/latest/getting_started/

### I don't have linux but I want to do this somehow!

As stated above, the plugin library only works in Linux. Bug the go team to
support your system!

* Or use a linux virtualbox, and mount this directory.
