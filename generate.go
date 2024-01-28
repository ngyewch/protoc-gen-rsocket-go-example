//go:generate protoc --proto_path . --go_out=. --go_opt=module=github.com/ngyewch/protoc-gen-rsocket-go-example service.proto service1.proto service2.proto
//go:generate protoc --proto_path . --rsocket-go_out=. --rsocket-go_opt=module=github.com/ngyewch/protoc-gen-rsocket-go-example service.proto
package main
