## Setup

Setup Go environment (OSX): 
1. Install Go via homebrew: `brew install go`
1. modify `~/.bash_profile` or `~/.zshrc` to ensure that `GOPATH`, `GOBIN`, and `GOROOT` are available as environment 
   variables: 
    ```
    export GOPATH="${HOME}/.go" 
    export GOBIN="$GOPATH/bin"
    export GOROOT="$(brew --prefix golang)/libexec"
    export GO111MODULE=on
    export PATH="$PATH:${GOROOT}/bin:$GOBIN"
    ```

Compiling the schema: 
* ensure that protocol buffer compiler (`protoc`) is installed. 
    * OSX via homebrew: `brew install protobuf`
* install these go packages:
    ``` 
    go get google.golang.org/grpc
    go get github.com/golang/protobuf/protoc-gen-go
    ```
* the protocol buffer schema is located at [grpc/api/meteorite-landing-service.proto](grpc/api/meteorite-landing-service.proto).
* change directory to [grpc/](grpc/) then run `protoc -I api/ api/meteorite-landing-service.proto --go_out=plugins=grpc:api`
* copy the generated _yudhistira.dev_ folder to the corresponding [grpc/client](grpc/client) and [grpc/server](grpc/server) directories.

Running the benchmark:
* Both `client` and `server` have their own respective golang module. `cd` inside and run `go install` to download the 
  necessary dependencies.
* Startup respective server ([grpc/server/grpc_server.go](grpc/server/grpc_server.go) or 
  [restful/server/http_server.go](restful/server/http_server.go))
* Change directory to the respective client directory ([grpc/client](grpc/client) or [restful/client](grpc/client)) and 
  run  `go test -bench=.`

## Some results

1 concurrent user: grpc is 4.5x faster
```
grpc:    1069319 ns /op
restful: 4814045 ns /op
```

10 concurrent users: grpc is 3.68x faster
```
grpc:	 3676586 ns/op
restful: 13565108 ns/op
```

100 concurrent users: grpc is 3.91x faster
```
grpc:     28657121 ns/op
restful: 112159437 ns/op
```

**NOTE**: take this result with grain of salt.

Runs on:
* `go version go1.14.1 darwin/amd64`
* `libprotoc 3.11.4`
* grpc and protobuf golang modules: 	
    * `github.com/golang/protobuf v1.4.0`
    * `google.golang.org/grpc v1.29.1`
  	* `google.golang.org/protobuf v1.23.0`

