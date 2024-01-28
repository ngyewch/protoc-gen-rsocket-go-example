module github.com/ngyewch/protoc-gen-rsocket-go-example

go 1.21

replace github.com/ngyewch/protoc-gen-rsocket-go => ../protoc-gen-rsocket-go

require (
	github.com/ngyewch/protoc-gen-rsocket-go v0.0.0-20240128084244-ff606372816c
	github.com/rsocket/rsocket-go v0.8.12
	github.com/samber/mo v1.11.0
	github.com/spf13/cobra v1.8.0
	google.golang.org/protobuf v1.32.0
)

require (
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jjeffcaii/reactor-go v0.5.5 // indirect
	github.com/panjf2000/ants/v2 v2.5.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	go.uber.org/atomic v1.7.0 // indirect
)
