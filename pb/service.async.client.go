package pb

import (
	"context"
	pb "github.com/ngyewch/protoc-gen-rsocket-go-example/pb/request/pb"
	pb1 "github.com/ngyewch/protoc-gen-rsocket-go-example/pb/response/pb"
	runtime "github.com/ngyewch/protoc-gen-rsocket-go/runtime"
	rsocketgo "github.com/rsocket/rsocket-go"
	mo "github.com/samber/mo"
	proto "google.golang.org/protobuf/proto"
)

type TestServiceClientAsync struct {
	selector uint64
	handler  runtime.ClientRequestResponseHandlerAsync
}

func NewTestServiceClientAsync(selector uint64, handler runtime.ClientRequestResponseHandlerAsync) *TestServiceClientAsync {
	return &TestServiceClientAsync{
		handler:  handler,
		selector: selector,
	}
}
func NewRSocketTestServiceClientAsync(selector uint64, rs rsocketgo.RSocket) *TestServiceClientAsync {
	return &TestServiceClientAsync{
		handler:  runtime.RSocketClientRequestResponseHandlerAsync(rs),
		selector: selector,
	}
}
func (c *TestServiceClientAsync) Test(ctx context.Context, req *pb.TestRequest) *mo.Future[*pb1.TestResponse] {
	return mo.NewFuture(func(resolve func(*pb1.TestResponse), reject func(error)) {
		runtime.HandleClientRequestResponseAsync(ctx, c.selector, "Test", req, c.handler).
			Catch(func(err error) ([]byte, error) {
				reject(err)
				return nil, err
			}).
			Then(func(rspBytes []byte) ([]byte, error) {
				var rsp pb1.TestResponse
				err := proto.Unmarshal(rspBytes, &rsp)
				if err != nil {
					return nil, err
				}
				resolve(&rsp)
				return rspBytes, nil
			})
	})
}

type TestService2ClientAsync struct {
	selector uint64
	handler  runtime.ClientRequestResponseHandlerAsync
}

func NewTestService2ClientAsync(selector uint64, handler runtime.ClientRequestResponseHandlerAsync) *TestService2ClientAsync {
	return &TestService2ClientAsync{
		handler:  handler,
		selector: selector,
	}
}
func NewRSocketTestService2ClientAsync(selector uint64, rs rsocketgo.RSocket) *TestService2ClientAsync {
	return &TestService2ClientAsync{
		handler:  runtime.RSocketClientRequestResponseHandlerAsync(rs),
		selector: selector,
	}
}
func (c *TestService2ClientAsync) Test(ctx context.Context, req *pb.TestRequest) *mo.Future[*pb1.TestResponse] {
	return mo.NewFuture(func(resolve func(*pb1.TestResponse), reject func(error)) {
		runtime.HandleClientRequestResponseAsync(ctx, c.selector, "Test", req, c.handler).
			Catch(func(err error) ([]byte, error) {
				reject(err)
				return nil, err
			}).
			Then(func(rspBytes []byte) ([]byte, error) {
				var rsp pb1.TestResponse
				err := proto.Unmarshal(rspBytes, &rsp)
				if err != nil {
					return nil, err
				}
				resolve(&rsp)
				return rspBytes, nil
			})
	})
}
