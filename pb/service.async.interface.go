package pb

import (
	"context"
	pb "github.com/ngyewch/protoc-gen-rsocket-go-example/pb/request/pb"
	pb1 "github.com/ngyewch/protoc-gen-rsocket-go-example/pb/response/pb"
	mo "github.com/samber/mo"
)

type TestServiceAsync interface {
	Test(context.Context, *pb.TestRequest) *mo.Future[*pb1.TestResponse]
}
type TestService2Async interface {
	Test(context.Context, *pb.TestRequest) *mo.Future[*pb1.TestResponse]
}
type TestListenerAsync interface {
	OnEvent(context.Context, *Event) *mo.Future[*Void]
}
