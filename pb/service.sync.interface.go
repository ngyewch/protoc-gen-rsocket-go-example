package pb

import (
	"context"
	pb "github.com/ngyewch/protoc-gen-rsocket-go-example/pb/request/pb"
	pb1 "github.com/ngyewch/protoc-gen-rsocket-go-example/pb/response/pb"
)

type TestService interface {
	Test(context.Context, *pb.TestRequest) (*pb1.TestResponse, error)
}
type TestService2 interface {
	Test(context.Context, *pb.TestRequest) (*pb1.TestResponse, error)
}
type TestListener interface {
	OnEvent(context.Context, *Event) (*Void, error)
}
